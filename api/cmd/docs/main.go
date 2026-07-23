// Command docs is the docs-site backend: documentation hub.
package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	foundationabuse "github.com/yueli-official/foundation/go/abuse"
	"github.com/yueli-official/foundation/go/abuse/turnstile"
	"github.com/yueli-official/foundation/go/authorization"
	authorizationpostgres "github.com/yueli-official/foundation/go/authorization/postgres"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"

	"platform/gokit/authsetup"
	"platform/gokit/observability"
	"platform/gokit/openapiexport"
	"platform/products/docs/api/internal/appconfig"
	"platform/products/docs/api/internal/catalog"
	"platform/products/docs/api/internal/dao"
	"platform/products/docs/api/internal/docsabuse"
	"platform/products/docs/api/internal/docsaudit"
	"platform/products/docs/api/internal/docsauthz"
	"platform/products/docs/api/internal/docsdiscovery"
	"platform/products/docs/api/internal/docssearch"
	"platform/products/docs/api/internal/docsurls"
	"platform/products/docs/api/internal/server"
)

func main() {
	ctx := gctx.New()
	shutdown, err := observability.StartFromEnvironment(ctx, "docs-api")
	if err != nil {
		panic(err)
	}
	defer observability.ShutdownWithTimeout(shutdown)

	// ── Catalog logic (DB access) ─────────────────────────────────────────────
	store := dao.NewPG(g.DB())
	discoveryModule, discoveryCache, err := docsdiscovery.New(store, appconfig.DiscoveryConfig(ctx))
	if err != nil {
		panic(err)
	}
	cat := catalog.New(store).
		WithAssets(appconfig.BuildAssetClient(ctx), appconfig.CoverCategory(ctx))
	var (
		abuseChallenge *foundationabuse.ChallengeDefinition
		abuseVerifiers map[foundationabuse.ChallengeKind]foundationabuse.ChallengeVerifier
	)
	if secret := g.Cfg().MustGet(ctx, "docs.abuse.turnstile.secret").String(); secret != "" && !openapiexport.Requested() {
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
	if openapiexport.Requested() {
		urlLifecycle, err := docsurls.NewMemory(appconfig.SiteURL(ctx), appconfig.DefaultLocale(ctx))
		if err != nil {
			panic(err)
		}
		cat.WithURLLifecycle(urlLifecycle)
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
		abuseModule, err := foundationabuse.NewMemory(abuseCatalog, foundationabuse.MemoryOptions{
			Secret:    []byte("docs-openapi-abuse-memory-secret"),
			Verifiers: abuseVerifiers,
		})
		if err != nil {
			panic(err)
		}
		authorizationService := docsauthz.New(authz)
		if err := authorizationService.SetAbuse(abuseModule); err != nil {
			panic(err)
		}
		jw := appconfig.LoadJWKS(ctx)
		verifier, err := authsetup.NewRemoteVerifier(authsetup.RemoteVerifierConfig{
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
			URLResolver: urlLifecycle.Resolver(),
		})
		if _, err := openapiexport.ExportIfRequested(s); err != nil {
			panic(err)
		}
		return
	}

	// ── Instance-local authorization ─────────────────────────────────────────
	authDB, err := appconfig.OpenAuthorizationDB(ctx)
	if err != nil {
		panic(err)
	}
	defer authDB.Close()
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
	verifier, err := authsetup.NewRemoteVerifier(authsetup.RemoteVerifierConfig{
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
		URLResolver: urlLifecycle.Resolver(),
	})
	if handled, err := openapiexport.ExportIfRequested(s); handled {
		if err != nil {
			panic(err)
		}
		return
	}
	g.Log().Info(ctx, "docs-service starting")
	s.Run()
}
