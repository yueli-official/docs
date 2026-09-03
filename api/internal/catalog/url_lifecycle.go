package catalog

import (
	"context"
	"database/sql"

	"github.com/yueli-official/docs/api/internal/dao"
	"github.com/yueli-official/docs/api/internal/docsurls"
)

func (s *Service) WithURLLifecycle(lifecycle *docsurls.Lifecycle) *Service {
	s.urls = lifecycle
	return s
}

func (s *Service) urlReconcileHook(collectionID, reason string) dao.TransactionHook {
	if s.urls == nil {
		return nil
	}
	return func(ctx context.Context, tx *sql.Tx) error {
		return s.urls.ReconcileCollection(ctx, tx, collectionID, reason)
	}
}

func (s *Service) urlInitializeHook(collectionID, reason string) dao.TransactionHook {
	if s.urls == nil {
		return nil
	}
	return func(ctx context.Context, tx *sql.Tx) error {
		return s.urls.InitializeCollection(ctx, tx, collectionID, reason)
	}
}
