package projectdocs

import (
	"context"
	"database/sql"
	"time"
)

type Source struct {
	ID              string     `json:"id"`
	CollectionID    string     `json:"collectionId"`
	CollectionSlug  string     `json:"collectionSlug"`
	CollectionTitle string     `json:"collectionTitle"`
	Repository      string     `json:"repository"`
	AssetName       string     `json:"assetName"`
	DefaultLocale   string     `json:"defaultLocale"`
	Mode            string     `json:"mode"`
	OwnerSub        string     `json:"ownerSub"`
	TokenCiphertext string     `json:"-"`
	Enabled         bool       `json:"enabled"`
	AutoCheck       bool       `json:"autoCheck"`
	Status          string     `json:"status"`
	LastReleaseID   int64      `json:"lastReleaseId"`
	LastAssetID     int64      `json:"lastAssetId"`
	LastTag         string     `json:"lastTag"`
	LastDigest      string     `json:"lastDigest"`
	LastError       string     `json:"lastError"`
	CheckedAt       *time.Time `json:"checkedAt"`
	NextCheckAt     time.Time  `json:"nextCheckAt"`
}
type Run struct {
	ID           string     `json:"id"`
	Status       string     `json:"status"`
	Stage        string     `json:"stage"`
	Tag          string     `json:"tag"`
	Digest       string     `json:"digest"`
	BatchID      string     `json:"batchId"`
	ErrorMessage string     `json:"errorMessage"`
	StartedAt    time.Time  `json:"startedAt"`
	CompletedAt  *time.Time `json:"completedAt"`
}

const sourceSelect = `SELECT s.id::text,s.collection_id::text,c.slug,c.title,s.repository,s.asset_name,s.default_locale,s.mode,s.owner_sub,s.token_ciphertext,s.enabled,s.status,s.last_release_id,s.last_asset_id,s.last_tag,s.last_digest,s.last_error,s.checked_at,s.next_check_at,s.auto_check FROM project_doc_sources s JOIN collections c ON c.id=s.collection_id`

type scanner interface{ Scan(...any) error }

func scanSource(row scanner) (Source, error) {
	var s Source
	err := row.Scan(&s.ID, &s.CollectionID, &s.CollectionSlug, &s.CollectionTitle, &s.Repository, &s.AssetName, &s.DefaultLocale, &s.Mode, &s.OwnerSub, &s.TokenCiphertext, &s.Enabled, &s.Status, &s.LastReleaseID, &s.LastAssetID, &s.LastTag, &s.LastDigest, &s.LastError, &s.CheckedAt, &s.NextCheckAt, &s.AutoCheck)
	return s, err
}

type Store struct{ DB *sql.DB }

func (s Store) Get(ctx context.Context, id string) (Source, error) {
	return scanSource(s.DB.QueryRowContext(ctx, sourceSelect+` WHERE s.id=$1`, id))
}
func (s Store) List(ctx context.Context, page, size int) ([]Source, int, error) {
	var total int
	if err := s.DB.QueryRowContext(ctx, `SELECT count(*) FROM project_doc_sources`).Scan(&total); err != nil {
		return nil, 0, err
	}
	rows, err := s.DB.QueryContext(ctx, sourceSelect+` ORDER BY s.created_at DESC LIMIT $1 OFFSET $2`, size, (page-1)*size)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()
	out := []Source{}
	for rows.Next() {
		item, e := scanSource(rows)
		if e != nil {
			return nil, 0, e
		}
		out = append(out, item)
	}
	return out, total, rows.Err()
}
func (s Store) Create(ctx context.Context, in Source) error {
	_, err := s.DB.ExecContext(ctx, `INSERT INTO project_doc_sources(id,collection_id,repository,asset_name,default_locale,mode,owner_sub,token_ciphertext,auto_check,status) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,CASE WHEN $9 THEN 'queued' ELSE 'idle' END)`, in.ID, in.CollectionID, in.Repository, in.AssetName, in.DefaultLocale, in.Mode, in.OwnerSub, in.TokenCiphertext, in.AutoCheck)
	return err
}
func (s Store) Runs(ctx context.Context, id string) ([]Run, error) {
	rows, err := s.DB.QueryContext(ctx, `SELECT id::text,status,stage,tag,digest,COALESCE(batch_id::text,''),error_message,started_at,completed_at FROM project_doc_source_runs WHERE source_id=$1 ORDER BY started_at DESC LIMIT 20`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	out := []Run{}
	for rows.Next() {
		var r Run
		if err = rows.Scan(&r.ID, &r.Status, &r.Stage, &r.Tag, &r.Digest, &r.BatchID, &r.ErrorMessage, &r.StartedAt, &r.CompletedAt); err != nil {
			return nil, err
		}
		out = append(out, r)
	}
	return out, rows.Err()
}
