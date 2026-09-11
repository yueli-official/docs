// Command docs is the docs-site backend: documentation hub.
package main

import (
	"context"
	"github.com/yueli-official/asset/referencesync"
	"github.com/yueli-official/docs/api/internal/assetreferences"
	"os"
	"strings"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	foundationabuse "github.com/yueli-official/foundation/go/abuse"
	"github.com/yueli-official/foundation/go/abuse/turnstile"
	"github.com/yueli-official/foundation/go/authorization"
	authorizationpostgres "github.com/yueli-official/foundation/go/authorization/postgres"
	"github.com/yueli-official/foundation/go/traffic"
	trafficpostgres "github.com/yueli-official/foundation/go/traffic/postgres"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"

	"github.com/yueli-official/docs/api/internal/appconfig"
	"github.com/yueli-official/docs/api/internal/catalog"
	"github.com/yueli-official/docs/api/internal/dao"
	"github.com/yueli-official/docs/api/internal/docsabuse"
	"github.com/yueli-official/docs/api/internal/docsanalytics"
	"github.com/yueli-official/docs/api/internal/docsaudit"
	"github.com/yueli-official/docs/api/internal/docsauthz"
	"github.com/yueli-official/docs/api/internal/docscomments"
	"github.com/yueli-official/docs/api/internal/docsdiscovery"
	"github.com/yueli-official/docs/api/internal/docssearch"
	"github.com/yueli-official/docs/api/internal/docstraffic"
	"github.com/yueli-official/docs/api/internal/docsurls"
	"github.com/yueli-official/docs/api/internal/identityclient"
	"github.com/yueli-official/docs/api/internal/runtime"
	"github.com/yueli-official/docs/api/internal/server"
)

func main() {
	if err := runtime.EnableEnvironmentConfig(); err != nil {
		panic(err)
	}
	ctx := gctx.New()
	if runtime.OpenAPIRequested() {
		exportOpenAPI(ctx)
		return
	}
	shutdown, err := runtime.StartTelemetry(ctx, "docs-api")
	if err != nil {
		panic(err)
	}
	defer runtime.ShutdownTelemetry(shutdown)

	// ── Catalog logic (DB access) ─────────────────────────────────────────────
	store := dao.NewPG(g.DB())
	discoveryModule, discoveryCache, err := docsdiscovery.New(store, appconfig.DiscoveryConfig(ctx))
	if err != nil {
		panic(err)
	}
	trafficDB, err := appconfig.OpenTrafficDB(ctx)
	if err != nil {
		panic(err)
	}
	defer trafficDB.Close()
	trafficCatalog, err := traffic.Compile(docstraffic.Definition(appconfig.TrafficTimeZone(ctx)))
	if err != nil {
		panic(err)
	}
	trafficModule, err := trafficpostgres.New(ctx, trafficCatalog, trafficpostgres.Options{
		DB: trafficDB, InstanceKey: "docs:" + appconfig.SiteSlug(ctx),
	})
	if err != nil {
		panic(err)
	}
	analyticsModule, err := docsanalytics.New(store, trafficModule, appconfig.TrafficTimeZone(ctx))
	if err != nil {
		panic(err)
	}
	trafficLocation, err := time.LoadLocation(appconfig.TrafficTimeZone(ctx))
	if err != nil {
		panic(err)
	}
	if strings.EqualFold(strings.TrimSpace(os.Getenv("DOCS_DEV_SEED")), "true") {
		if err := docsanalytics.SeedLocal(ctx, analyticsModule, time.Now().In(trafficLocation)); err != nil {
			panic(err)
		}
	}
	if err := store.PruneAnalyticsTrafficSourceReceipts(ctx, time.Now().In(trafficLocation).AddDate(0, 0, -60)); err != nil {
		panic(err)
	}
	cat := catalog.New(store).
		WithAssets(appconfig.BuildAssetClient(ctx), appconfig.CoverCategory(ctx))
	var (
		abuseChallenge *foundationabuse.ChallengeDefinition
		abuseVerifiers map[foundationabuse.ChallengeKind]foundationabuse.ChallengeVerifier
	)
	if secret := g.Cfg().MustGet(ctx, "docs.abuse.turnstile.secret").String(); secret != "" {
		hostnames := g.Cfg().MustGet(ctx, "docs.abuse.turnstile.hostnames").Strings()
		if len(hostnames) == 0 {
			panic("docs.abuse.turnstile.hostnames is required when Turnstile is enabled")
		}
		challengeVerifier, err := turnstile.New(turnstile.Options{
			Secret:   secret,
			Endpoint: g.Cfg().MustGet(ctx, "docs.abuse.turnstile.endpoint").String(),
		})
		if err != nil {
			panic(err)
		}
		abuseChallenge = &foundationabuse.ChallengeDefinition{
			Kind: "turnstile", ExpectedAction: "docs-author-application",
			AllowedHosts: hostnames,
		}
		abuseVerifiers = map[foundationabuse.ChallengeKind]foundationabuse.ChallengeVerifier{
			"turnstile": challengeVerifier,
		}
	}
	abuseCatalog := foundationabuse.MustCompile(docsabuse.Definition(docsabuse.Policy{
		Challenge: abuseChallenge,
	}))

	// ── Instance-local authorization ─────────────────────────────────────────
	authDB, err := appconfig.OpenAuthorizationDB(ctx)
	if err != nil {
		panic(err)
	}
	defer authDB.Close()
	commentsModule := docscomments.New(docscomments.NewPostgres(authDB))
	identityProfiles := identityclient.NewHTTP(appconfig.IdentityBaseURL(ctx))
	auditJournal, err := docsaudit.New(ctx, authDB, appconfig.SiteSlug(ctx))
	if err != nil {
		panic(err)
	}
	cat.WithAudit(auditJournal)
	searchIndex, err := docssearch.NewPostgres(ctx, authDB, appconfig.SiteSlug(ctx))
	if err != nil {
		panic(err)
	}
	if err := searchIndex.Reconcile(ctx, authDB); err != nil {
		panic(err)
	}
	cat.WithSearch(searchIndex)
	definition, err := authorization.Compile(docsauthz.Definition())
	if err != nil {
		panic(err)
	}
	bootstrapSubs := appconfig.BootstrapAdministratorSubs(ctx)
	protected := make([]authorization.SubjectRef, 0, len(bootstrapSubs))
	for _, sub := range bootstrapSubs {
		if sub != "" {
			protected = append(protected, authorization.SubjectRef{Kind: authorization.SubjectUser, ID: sub})
		}
	}
	authz, err := authorizationpostgres.New(ctx, definition, authorizationpostgres.Options{
		DB: authDB, InstanceKey: "docs:" + appconfig.SiteSlug(ctx),
		Memory: authorization.MemoryOptions{
			RootScopeID:       docsauthz.RootScopeID,
			ProtectedSubjects: protected,
			Constraints:       docsauthz.ConstraintEvaluators(),
			Predicates:        docsauthz.PredicateEvaluators(),
		},
	})
	if err != nil {
		panic(err)
	}
	abuseModule, err := foundationabuse.NewPostgres(ctx, abuseCatalog, foundationabuse.PostgresOptions{
		DB: authDB, InstanceKey: "docs:" + appconfig.SiteSlug(ctx),
		Verifiers: abuseVerifiers,
	})
	if err != nil {
		panic(err)
	}
	authorizationService := docsauthz.NewWithDB(authz, authDB)
	if err := authorizationService.SetAbuse(abuseModule); err != nil {
		panic(err)
	}
	if authz.InstanceWasCreated() {
		if len(protected) == 0 {
			panic("docs authorization bootstrap requires at least one administrator subject")
		}
		if err := docsauthz.SyncResourceScopes(ctx, authDB, authz); err != nil {
			panic(err)
		}
	}
	urlLifecycle, err := docsurls.NewPostgres(
		ctx,
		authDB,
		"docs:"+appconfig.SiteSlug(ctx),
		appconfig.SiteURL(ctx),
		appconfig.DefaultLocale(ctx),
	)
	if err != nil {
		panic(err)
	}
	if err := urlLifecycle.ReconcileAll(ctx, authDB); err != nil {
		panic(err)
	}
	cat.WithURLLifecycle(urlLifecycle)

	// ── JWT verifier (IdP JWKS, lazy) ────────────────────────────────────────
	jw := appconfig.LoadJWKS(ctx)
	verifier, err := runtime.NewRemoteVerifier(runtime.RemoteVerifierConfig{
		JWKSURL: jw.URL, Issuer: jw.Issuer, Audience: jw.Audience,
		AllowLoopbackHTTP: jw.AllowLoopbackHTTP,
	})
	if err != nil {
		panic(err)
	}

	s := g.Server()
	server.Configure(s, server.Deps{
		Verifier: verifier, Catalog: cat, Authorization: authorizationService,
		Discovery: discoveryModule, DiscoveryCache: discoveryCache,
		URLResolver: urlLifecycle.Resolver(), Analytics: analyticsModule,
		Comments: commentsModule,
		Identity: identityProfiles,
	})
	refSecret := g.Cfg().MustGet(ctx, "docs.assetService.clientSecret").String()
	if refSecret != "" {
		refClient := &referencesync.Client{BaseURL: g.Cfg().MustGet(ctx, "docs.assetService.baseUrl").String(), Token: referencesync.ClientCredentials(
			strings.TrimRight(appconfig.IdentityBaseURL(ctx), "/")+"/oauth2/token", g.Cfg().MustGet(ctx, "docs.assetService.clientId", "docs-asset-svc").String(), refSecret)}
		go referencesync.Run(ctx, authDB, "docs:asset-references", assetreferences.Source(appconfig.SiteURL(ctx), g.Cfg().MustGet(ctx, "docs.assetService.publicOrigin").String()), refClient, func(err error) { g.Log().Warning(ctx, "asset reference reconciliation:", err) })
	}
	g.Log().Info(ctx, "docs-service starting")
	s.Run()
}

func exportOpenAPI(ctx context.Context) {
	discoveryModule, discoveryCache, err := docsdiscovery.New(nil, docsdiscovery.Config{
		Origin:        "https://docs.example.test",
		Name:          "Docs",
		Description:   "OpenAPI export",
		DefaultLocale: "en",
		TTL:           5 * time.Minute,
		Clock:         time.Now,
	})
	if err != nil {
		panic(err)
	}
	trafficCatalog, err := traffic.Compile(docstraffic.Definition("UTC"))
	if err != nil {
		panic(err)
	}
	trafficModule, err := traffic.NewMemory(trafficCatalog, traffic.MemoryOptions{
		Clock: time.Now, Secret: []byte("docs-openapi-traffic-secret-32-bytes"),
	})
	if err != nil {
		panic(err)
	}
	analyticsModule, err := docsanalytics.New(dao.NewPG(nil), trafficModule, "UTC")
	if err != nil {
		panic(err)
	}
	urlLifecycle, err := docsurls.NewMemory("https://docs.example.test", "en")
	if err != nil {
		panic(err)
	}
	cat := catalog.New(nil).WithURLLifecycle(urlLifecycle)
	definition, err := authorization.Compile(docsauthz.Definition())
	if err != nil {
		panic(err)
	}
	authz, err := authorization.NewMemory(definition, authorization.MemoryOptions{
		RootScopeID:       docsauthz.RootScopeID,
		ProtectedSubjects: []authorization.SubjectRef{{Kind: authorization.SubjectUser, ID: "openapi-export-admin"}},
		Constraints:       docsauthz.ConstraintEvaluators(),
		Predicates:        docsauthz.PredicateEvaluators(),
	})
	if err != nil {
		panic(err)
	}
	abuseCatalog := foundationabuse.MustCompile(docsabuse.Definition(docsabuse.Policy{}))
	abuseModule, err := foundationabuse.NewMemory(abuseCatalog, foundationabuse.MemoryOptions{
		Secret: []byte("docs-openapi-abuse-memory-secret"),
	})
	if err != nil {
		panic(err)
	}
	authorizationService := docsauthz.New(authz)
	if err := authorizationService.SetAbuse(abuseModule); err != nil {
		panic(err)
	}
	s := g.Server()
	server.Configure(s, server.Deps{
		Catalog: cat, Authorization: authorizationService,
		Discovery: discoveryModule, DiscoveryCache: discoveryCache,
		URLResolver: urlLifecycle.Resolver(), Analytics: analyticsModule,
		Comments: docscomments.New(docscomments.NewMemory()),
	})
	handled, err := runtime.ExportOpenAPIIfRequested(s)
	if err != nil {
		panic(err)
	}
	if !handled {
		panic("DOCS_OPENAPI_OUTPUT is required")
	}
}
