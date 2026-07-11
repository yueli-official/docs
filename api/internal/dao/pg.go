// Package dao provides database-access objects for the docs service.
package dao

import "github.com/gogf/gf/v2/database/gdb"

// PG wraps a GoFrame database handle for the docs service.
type PG struct{ db gdb.DB }

// NewPG returns a new PG dao wrapping the given database handle.
func NewPG(db gdb.DB) *PG { return &PG{db: db} }
