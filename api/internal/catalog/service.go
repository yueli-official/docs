// Package catalog implements the docs-site document catalog logic.
package catalog

import (
	"context"
	"github.com/yueli-official/docs/api/internal/assetclient"
	"github.com/yueli-official/docs/api/internal/dao"
	"github.com/yueli-official/docs/api/internal/docsaudit"
	"github.com/yueli-official/docs/api/internal/docssearch"
	"github.com/yueli-official/docs/api/internal/docsurls"
)

// Service holds the catalog business logic.
type Service struct {
	importGuard   func(context.Context, string) error
	dao           *dao.PG
	asset         assetclient.Client
	coverCategory string
	urls          *docsurls.Lifecycle
	audit         *docsaudit.Journal
	search        *docssearch.Index
}

// New returns a new catalog Service backed by the given dao. Asset integration
// is optional so existing tests and read-only deployments can build the catalog
// without an asset service.
func New(d *dao.PG) *Service { return &Service{dao: d, coverCategory: "docs-collection-cover"} }

func (s *Service) WithAudit(journal *docsaudit.Journal) *Service {
	s.audit = journal
	return s
}

func (s *Service) WithSearch(index *docssearch.Index) *Service {
	s.search = index
	return s
}

// WithAssets wires the asset service for collection cover uploads.
func (s *Service) WithAssets(asset assetclient.Client, coverCategory string) *Service {
	s.asset = asset
	if coverCategory != "" {
		s.coverCategory = coverCategory
	}
	return s
}

func (s *Service) WithImportGuard(guard func(context.Context, string) error) *Service {
	s.importGuard = guard
	return s
}

type importCommitCheckKey struct{}

// WithImportCommitCheck lets a background source verify it is still enabled at
// commit time. It supplements, and never replaces, the user's authorization.
func WithImportCommitCheck(ctx context.Context, check func(context.Context) error) context.Context {
	return context.WithValue(ctx, importCommitCheckKey{}, check)
}
