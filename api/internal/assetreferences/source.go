// Package assetreferences projects current document usage into Asset references.
package assetreferences

import (
	"context"
	"database/sql"
	"github.com/yueli-official/asset/referencesync"
	"strings"
)

// MediaKeys delegates media URL parsing to the Asset SDK.
func MediaKeys(content string, origins ...string) []string {
	return referencesync.MediaKeys(content, origins...)
}

func Source(origin, assetOrigin string) func(context.Context, *sql.Tx) ([]referencesync.Snapshot, error) {
	return func(ctx context.Context, tx *sql.Tx) ([]referencesync.Snapshot, error) {
		documents := referencesync.Snapshot{RefType: "document", References: []referencesync.Reference{}}
		covers := referencesync.Snapshot{RefType: "collection-cover", References: []referencesync.Reference{}}
		rows, err := tx.QueryContext(ctx, `SELECT d.id::text,d.title,d.content FROM docs d JOIN collections c ON c.id=d.collection_id WHERE d.deleted_at IS NULL ORDER BY d.id`)
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			var id, title, content string
			if err = rows.Scan(&id, &title, &content); err != nil {
				rows.Close()
				return nil, err
			}
			for _, key := range MediaKeys(content, origin, assetOrigin) {
				documents.References = append(documents.References, referencesync.Reference{MediaKey: key, RefID: id, RefLabel: title, RefURL: strings.TrimRight(origin, "/") + "/manage/docs/" + id})
			}
		}
		err = rows.Err()
		rows.Close()
		if err != nil {
			return nil, err
		}
		rows, err = tx.QueryContext(ctx, `SELECT id::text,title,slug,cover_asset_id::text FROM collections WHERE cover_asset_id IS NOT NULL AND cover_asset_id::text<>'' ORDER BY id`)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		for rows.Next() {
			var id, title, slug, asset string
			if err = rows.Scan(&id, &title, &slug, &asset); err != nil {
				return nil, err
			}
			covers.References = append(covers.References, referencesync.Reference{AssetID: asset, RefID: id, RefLabel: title, RefURL: strings.TrimRight(origin, "/") + "/" + slug})
		}
		if err = rows.Err(); err != nil {
			return nil, err
		}
		return []referencesync.Snapshot{documents, covers}, nil
	}
}
