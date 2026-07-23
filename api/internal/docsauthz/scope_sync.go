package docsauthz

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/yueli-official/foundation/go/authorization"
)

// SyncResourceScopes reconciles pre-existing catalog resources into the
// authorization scope tree. It is idempotent and is intended for bootstrap and
// repair, while normal creates register their scope immediately.
func SyncResourceScopes(
	ctx context.Context,
	db *sql.DB,
	runtime authorization.ResourceScopeRegistry,
) error {
	rows, err := db.QueryContext(ctx, `SELECT id::text FROM collections ORDER BY id`)
	if err != nil {
		return fmt.Errorf("list collections for authorization scope sync: %w", err)
	}
	var collectionIDs []string
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			_ = rows.Close()
			return fmt.Errorf("scan collection for authorization scope sync: %w", err)
		}
		collectionIDs = append(collectionIDs, id)
	}
	if err := rows.Close(); err != nil {
		return fmt.Errorf("close collection scope sync rows: %w", err)
	}
	if err := rows.Err(); err != nil {
		return fmt.Errorf("iterate collections for authorization scope sync: %w", err)
	}
	for _, id := range collectionIDs {
		if err := ensureScope(ctx, runtime, authorization.RegisterScopeCommand{
			ID: CollectionScopeID(id), Type: ScopeCollection, ParentID: RootScopeID,
		}); err != nil {
			return err
		}
	}

	rows, err = db.QueryContext(ctx, `SELECT id::text, collection_id::text FROM docs ORDER BY id`)
	if err != nil {
		return fmt.Errorf("list documents for authorization scope sync: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var id, collectionID string
		if err := rows.Scan(&id, &collectionID); err != nil {
			return fmt.Errorf("scan document for authorization scope sync: %w", err)
		}
		if err := ensureScope(ctx, runtime, authorization.RegisterScopeCommand{
			ID: DocumentScopeID(id), Type: ScopeDocument, ParentID: CollectionScopeID(collectionID),
		}); err != nil {
			return err
		}
	}
	if err := rows.Err(); err != nil {
		return fmt.Errorf("iterate documents for authorization scope sync: %w", err)
	}
	return nil
}

func ensureScope(
	ctx context.Context,
	runtime authorization.ResourceScopeRegistry,
	command authorization.RegisterScopeCommand,
) error {
	_, err := runtime.RegisterScope(ctx, command)
	if err == nil || authorization.Is(err, authorization.ErrorConflict) {
		return nil
	}
	return fmt.Errorf("create authorization scope %q: %w", command.ID, err)
}
