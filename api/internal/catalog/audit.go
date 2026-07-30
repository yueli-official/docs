package catalog

import (
	"context"
	"crypto/sha256"
	"encoding/hex"

	"github.com/google/uuid"
	"github.com/yueli-official/foundation/go/audit"

	"platform/products/docs/api/internal/dao"
	"platform/products/docs/api/internal/docsaudit"
	"platform/products/docs/api/internal/model"
)

func (s *Service) documentMutationHook(
	ctx context.Context,
	action docsaudit.Action,
	doc *model.Doc,
	reason string,
) dao.TransactionHook {
	var auditHook dao.TransactionHook
	if s.audit != nil && action != "" {
		auditHook = s.audit.Hook(
			ctx, action, uuid.NewString(),
			audit.Target{Type: "docs.document", ID: doc.ID},
			docsaudit.Evidence{Digest: documentDigest(doc)}, "",
		)
	}
	return dao.ComposeTransactionHooks(
		s.urlReconcileHook(doc.CollectionID, reason),
		auditHook,
		s.searchHook(doc.ID),
	)
}

func (s *Service) searchHook(id string) dao.TransactionHook {
	if s.search == nil {
		return nil
	}
	return s.search.SubtreeHook(id)
}

func (s *Service) importMutationHook(
	ctx context.Context,
	action docsaudit.Action,
	batch *model.ImportBatch,
	count int,
	actor string,
	reason string,
) dao.TransactionHook {
	var auditHook dao.TransactionHook
	if s.audit != nil {
		auditHook = s.audit.Hook(
			ctx, action, uuid.NewString(),
			audit.Target{Type: "docs.import", ID: batch.ID},
			docsaudit.Evidence{Count: uint64(count)}, actor,
		)
	}
	return dao.ComposeTransactionHooks(
		s.urlReconcileHook(batch.CollectionID, reason),
		auditHook,
		s.searchCollectionHook(batch.CollectionID),
	)
}

func (s *Service) searchCollectionHook(collectionID string) dao.TransactionHook {
	if s.search == nil {
		return nil
	}
	return s.search.CollectionHook(collectionID)
}

func documentDigest(doc *model.Doc) string {
	raw := []byte(
		doc.ID + "\x00" + doc.CollectionID + "\x00" + doc.VersionID + "\x00" +
			doc.ParentID + "\x00" + doc.Slug + "\x00" + doc.Title + "\x00" +
			doc.Content + "\x00" + doc.Status + "\x00" + doc.Locale,
	)
	sum := sha256.Sum256(raw)
	return hex.EncodeToString(sum[:])
}
