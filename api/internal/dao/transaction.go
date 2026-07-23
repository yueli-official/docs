package dao

import (
	"context"
	"database/sql"

	"github.com/gogf/gf/v2/database/gdb"
)

// TransactionHook updates a co-located Foundation module using the same
// caller-owned product transaction.
type TransactionHook func(context.Context, *sql.Tx) error

func runTransactionHook(ctx context.Context, tx gdb.TX, hook TransactionHook) error {
	if hook == nil {
		return nil
	}
	return hook(ctx, tx.GetSqlTX())
}
