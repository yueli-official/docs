package projectdocs

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/yueli-official/docs/api/internal/catalog"
	"github.com/yueli-official/docs/api/internal/docsauthz"
	"github.com/yueli-official/docs/api/internal/docserr"
	auth "github.com/yueli-official/foundation/go/auth"
	"github.com/yueli-official/foundation/go/authorization"
	"github.com/yueli-official/foundation/go/identifier"
)

type ReleaseClient interface {
	Latest(context.Context, string, string) (Release, ReleaseAsset, error)
	Download(context.Context, ReleaseAsset) ([]byte, string, error)
}
type Service struct {
	Store         Store
	catalog       *catalog.Service
	authorization *docsauthz.Service
	verifier      *auth.PersonalTokenVerifier
	vault         *Vault
	github        ReleaseClient
	wake          chan struct{}
}

func New(db *sql.DB, c *catalog.Service, a *docsauthz.Service, v *auth.PersonalTokenVerifier, vault *Vault) *Service {
	return &Service{Store: Store{DB: db}, catalog: c, authorization: a, verifier: v, vault: vault, github: NewGitHub(""), wake: make(chan struct{}, 1)}
}

func (s *Service) WithGitHubToken(token string) *Service { s.github = NewGitHub(token); return s }

type CreateInput struct {
	Repository      string
	AssetName       string
	Collection      string
	DefaultLocale   string
	Mode            string
	Token           string
	ExclusiveMirror bool
	AutoCheck       bool
}

func (s *Service) credential(ctx context.Context, token, owner string) (context.Context, error) {
	if s.verifier == nil || !strings.HasPrefix(token, "pat_") {
		return nil, docserr.Forbidden()
	}
	p, err := s.verifier.Verify(ctx, token)
	if err != nil || p == nil || p.Subject != owner || !p.IsPersonalToken() {
		return nil, docserr.Forbidden()
	}
	userCtx := auth.NewContext(ctx, p)
	d, err := s.authorization.Decide(userCtx, docsauthz.CapabilityImportManage, docsauthz.RootScopeID, authorization.ResourceFacts{})
	if err != nil || !d.Allowed {
		return nil, docserr.Forbidden()
	}
	return userCtx, nil
}
func (s *Service) Notify() {
	select {
	case s.wake <- struct{}{}:
	default:
	}
}
func (s *Service) Create(ctx context.Context, owner string, in CreateInput) (Source, error) {
	repo, err := Repository(in.Repository)
	if err != nil {
		return Source{}, docserr.InvalidInput("github_repository_invalid")
	}
	if in.AssetName == "" {
		in.AssetName = "docs.zip"
	}
	if len(in.AssetName) > 200 || strings.ContainsAny(in.AssetName, "/\\") || !strings.HasSuffix(strings.ToLower(in.AssetName), ".zip") {
		return Source{}, docserr.InvalidInput("github_asset_invalid")
	}
	if in.Mode == "" {
		in.Mode = "upsert"
	}
	if in.Mode != "upsert" && (in.Mode != "replace-version" || !in.ExclusiveMirror) {
		return Source{}, docserr.InvalidInput("exclusive_mirror_required")
	}
	if in.DefaultLocale == "" || len(in.DefaultLocale) > 32 {
		return Source{}, docserr.InvalidInput("default_locale_required")
	}
	if _, err = s.credential(ctx, in.Token, owner); err != nil {
		return Source{}, err
	}
	collection, err := s.catalog.GetCollectionBySlug(ctx, in.Collection)
	if err != nil {
		return Source{}, err
	}
	id := identifier.MustNew().String()
	sealed, err := s.vault.Seal(id, in.Token)
	if err != nil {
		return Source{}, err
	}
	item := Source{ID: id, CollectionID: collection.ID, Repository: repo, AssetName: in.AssetName, DefaultLocale: in.DefaultLocale, Mode: in.Mode, OwnerSub: owner, TokenCiphertext: sealed, AutoCheck: in.AutoCheck}
	if err = s.Store.Create(ctx, item); err != nil {
		// Map uniqueness without exposing driver messages or credential parameters.
		var state interface{ SQLState() string }
		if errors.As(err, &state) && state.SQLState() == "23505" {
			return Source{}, docserr.InvalidInput("collection_already_has_source")
		}
		return Source{}, errors.New("could not save documentation source")
	}
	s.Notify()
	return s.Store.Get(ctx, id)
}
func (s *Service) Get(ctx context.Context, id string) (Source, []Run, error) {
	source, err := s.Store.Get(ctx, id)
	if errors.Is(err, sql.ErrNoRows) {
		return Source{}, nil, docserr.NotFound(id)
	}
	if err != nil {
		return Source{}, nil, err
	}
	runs, err := s.Store.Runs(ctx, id)
	return source, runs, err
}
func (s *Service) Queue(ctx context.Context, id string) error {
	result, err := s.Store.DB.ExecContext(ctx, `UPDATE project_doc_sources SET next_check_at=NOW(),status=CASE WHEN status='running' THEN status ELSE 'queued' END,updated_at=NOW() WHERE id=$1 AND enabled`, id)
	if err != nil {
		return err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return docserr.InvalidInput("source_paused_or_missing")
	}
	s.Notify()
	return nil
}
func (s *Service) SetEnabled(ctx context.Context, id, owner, token string, enabled bool) error {
	return s.Update(ctx, id, owner, token, &enabled, nil)
}
func (s *Service) Update(ctx context.Context, id, owner, token string, enabled, autoCheck *bool) error {
	source, err := s.Store.Get(ctx, id)
	if errors.Is(err, sql.ErrNoRows) {
		return docserr.NotFound(id)
	}
	if err != nil {
		return err
	}
	sealed := source.TokenCiphertext
	if token != "" {
		if _, err = s.credential(ctx, token, owner); err != nil {
			return err
		}
		sealed, err = s.vault.Seal(id, token)
		if err != nil {
			return err
		}
	} else {
		owner = source.OwnerSub
	}
	if (enabled != nil && *enabled) || token != "" {
		credential, e := s.vault.Open(id, sealed)
		if e != nil {
			return docserr.Forbidden()
		}
		if _, e = s.credential(ctx, credential, owner); e != nil {
			return e
		}
	}
	_, err = s.Store.DB.ExecContext(ctx, `UPDATE project_doc_sources SET enabled=COALESCE($2,enabled),owner_sub=$3,token_ciphertext=$4,auto_check=COALESCE($5,auto_check),status=CASE WHEN $2::boolean IS NULL THEN status WHEN NOT $2 THEN 'paused' WHEN COALESCE($5,auto_check) THEN 'queued' ELSE 'idle' END,next_check_at=CASE WHEN $2::boolean=true OR $5::boolean=true THEN NOW() ELSE next_check_at END,updated_at=NOW() WHERE id=$1`, id, enabled, owner, sealed, autoCheck)
	if err == nil {
		s.Notify()
	}
	return err
}
func (s *Service) Delete(ctx context.Context, id string) error {
	result, err := s.Store.DB.ExecContext(ctx, `DELETE FROM project_doc_sources WHERE id=$1`, id)
	if err != nil {
		return err
	}
	n, err := result.RowsAffected()
	if n == 0 && err == nil {
		return docserr.NotFound(id)
	}
	return err
}
