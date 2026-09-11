package projectdocs

import (
	"context"
	"database/sql/driver"
	"errors"
	"time"

	"github.com/yueli-official/docs/api/internal/catalog"
	"github.com/yueli-official/foundation/go/identifier"
)

// Run polls durable due times. A connection-scoped advisory lock fences workers
// across processes and releases automatically if a process exits.
func (s *Service) Run(ctx context.Context) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()
	for {
		s.runDue(ctx)
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
		case <-s.wake:
		}
	}
}
func (s *Service) runDue(ctx context.Context) {
	rows, err := s.Store.DB.QueryContext(ctx, `SELECT id::text FROM project_doc_sources WHERE enabled AND (auto_check OR status IN ('queued','running')) AND next_check_at<=NOW() ORDER BY next_check_at LIMIT 10`)
	if err != nil {
		return
	}
	var ids []string
	for rows.Next() {
		var id string
		if rows.Scan(&id) == nil {
			ids = append(ids, id)
		}
	}
	_ = rows.Close()
	for _, id := range ids {
		if ctx.Err() != nil {
			return
		}
		s.check(ctx, id)
	}
}
func (s *Service) check(parent context.Context, id string) {
	ctx, cancel := context.WithTimeout(parent, 15*time.Minute)
	defer cancel()
	conn, err := s.Store.DB.Conn(ctx)
	if err != nil {
		return
	}
	defer conn.Close()
	var locked bool
	if err = conn.QueryRowContext(ctx, `SELECT pg_try_advisory_lock(hashtextextended('docs-project-source:' || $1,0))`, id).Scan(&locked); err != nil || !locked {
		return
	}
	defer func() {
		cleanup, c := context.WithTimeout(context.Background(), 5*time.Second)
		defer c()
		if _, e := conn.ExecContext(cleanup, `SELECT pg_advisory_unlock(hashtextextended('docs-project-source:' || $1,0))`, id); e != nil {
			_ = conn.Raw(func(any) error { return driver.ErrBadConn })
		}
	}()
	source, err := s.Store.Get(ctx, id)
	if err != nil || !source.Enabled || source.NextCheckAt.After(time.Now()) || (!source.AutoCheck && source.Status != "queued" && source.Status != "running") {
		return
	}
	// An old running record now has no live lock holder. Recover a committed
	// batch, otherwise record the interruption and retry from a fresh preflight.
	if err = s.recoverRun(ctx, source); err != nil {
		return
	}
	source, err = s.Store.Get(ctx, id)
	if err != nil {
		return
	}
	runID := identifier.MustNew().String()
	if _, err = s.Store.DB.ExecContext(ctx, `INSERT INTO project_doc_source_runs(id,source_id,status) VALUES($1,$2,'running')`, runID, id); err != nil {
		return
	}
	_, err = s.Store.DB.ExecContext(ctx, `UPDATE project_doc_sources SET status='running',checked_at=NOW(),next_check_at=NOW()+INTERVAL '15 minutes' WHERE id=$1`, id)
	if err != nil {
		return
	}
	stage := "authorization"
	var retryAt *time.Time
	finishFailure := func(message string, pause bool) {
		cleanup, c := context.WithTimeout(context.Background(), 10*time.Second)
		defer c()
		_, _ = s.Store.DB.ExecContext(cleanup, `UPDATE project_doc_source_runs SET status='failed',stage=$2,error_message=$3,completed_at=NOW() WHERE id=$1`, runID, stage, message)
		_, _ = s.Store.DB.ExecContext(cleanup, `UPDATE project_doc_sources SET status=CASE WHEN $3 OR NOT enabled THEN 'paused' ELSE 'failed' END,enabled=enabled AND NOT $3,last_error=$2,updated_at=NOW(),next_check_at=NOW()+INTERVAL '15 minutes' WHERE id=$1`, id, message, pause)
		if retryAt != nil {
			_, _ = s.Store.DB.ExecContext(cleanup, `UPDATE project_doc_sources SET next_check_at=GREATEST(next_check_at,$2) WHERE id=$1`, id, *retryAt)
		}
	}
	token, err := s.vault.Open(id, source.TokenCiphertext)
	if err != nil {
		finishFailure("同步凭证无法读取，请更换令牌。", true)
		return
	}
	userCtx, err := s.credential(ctx, token, source.OwnerSub)
	if err != nil {
		finishFailure("令牌或导入权限已失效，请更换令牌后恢复。", true)
		return
	}
	stage = "checking"
	release, asset, err := s.github.Latest(ctx, source.Repository, source.AssetName)
	if err != nil {
		var limited *RateLimitError
		if errors.As(err, &limited) {
			retryAt = &limited.RetryAt
		}
		finishFailure(err.Error(), false)
		return
	}
	_, err = s.Store.DB.ExecContext(ctx, `UPDATE project_doc_source_runs SET release_id=$2,asset_id=$3,tag=$4,stage='downloading' WHERE id=$1`, runID, release.ID, asset.ID, release.Tag)
	if err != nil {
		return
	}
	stage = "downloading"
	if source.LastReleaseID == release.ID && source.LastAssetID == asset.ID && source.LastDigest != "" && (asset.Digest == "" || asset.Digest == source.LastDigest) {
		s.complete(ctx, source, runID, release, asset, source.LastDigest, "skipped")
		return
	}
	data, digest, err := s.github.Download(ctx, asset)
	if err != nil {
		var limited *RateLimitError
		if errors.As(err, &limited) {
			retryAt = &limited.RetryAt
		}
		finishFailure(err.Error(), false)
		return
	}
	if digest == source.LastDigest {
		s.complete(ctx, source, runID, release, asset, digest, "skipped")
		return
	}
	_, err = s.Store.DB.ExecContext(ctx, `UPDATE project_doc_source_runs SET digest=$2,stage='preflight' WHERE id=$1`, runID, digest)
	if err != nil {
		return
	}
	stage = "preflight"
	batch, summary, err := s.catalog.PreflightImport(userCtx, catalog.ImportUploadInput{Filename: source.AssetName, Data: data, Bearer: token, Author: source.OwnerSub, Collection: source.CollectionSlug, DefaultLocale: source.DefaultLocale, Mode: source.Mode})
	if err != nil {
		finishFailure("文档包预检失败，请检查 ZIP 格式与目标文档集。", false)
		return
	}
	_, err = s.Store.DB.ExecContext(ctx, `UPDATE project_doc_source_runs SET batch_id=$2,stage='confirming' WHERE id=$1`, runID, batch.ID)
	if err != nil {
		return
	}
	if summary.Blocking {
		finishFailure("文档包存在阻塞问题，请打开导入报告。", false)
		return
	}
	stage = "confirming"
	userCtx = catalog.WithImportCommitCheck(userCtx, func(checkCtx context.Context) error {
		latest, e := s.Store.Get(checkCtx, id)
		if e != nil || !latest.Enabled || latest.TokenCiphertext != source.TokenCiphertext {
			return errors.New("documentation source changed or paused")
		}
		return nil
	})
	_, _, err = s.catalog.ConfirmImport(userCtx, batch.ID, token, source.OwnerSub)
	if err != nil {
		finishFailure("导入未完成，请检查令牌权限、图片处理和导入报告。", false)
		return
	}
	if err = s.syncScopes(userCtx, batch.ID, source.CollectionID); err != nil {
		finishFailure("文档已写入，权限索引待恢复。", false)
		return
	}
	s.complete(ctx, source, runID, release, asset, digest, "completed")
}

func (s *Service) syncScopes(ctx context.Context, batchID, collectionID string) error {
	_, items, err := s.catalog.GetImport(ctx, batchID)
	if err != nil {
		return err
	}
	if err = s.authorization.EnsureCollectionScope(ctx, collectionID); err != nil {
		return err
	}
	for _, item := range items {
		if item.TargetDocID != "" {
			if err = s.authorization.EnsureDocumentScope(ctx, item.TargetDocID, collectionID); err != nil {
				return err
			}
		}
	}
	return nil
}
func (s *Service) complete(ctx context.Context, source Source, runID string, release Release, asset ReleaseAsset, digest, status string) {
	tx, err := s.Store.DB.BeginTx(ctx, nil)
	if err != nil {
		return
	}
	defer tx.Rollback()
	_, err = tx.ExecContext(ctx, `UPDATE project_doc_source_runs SET status=$2,stage='done',digest=$3,completed_at=NOW(),error_message='' WHERE id=$1`, runID, status, digest)
	if err != nil {
		return
	}
	_, err = tx.ExecContext(ctx, `UPDATE project_doc_sources SET status=CASE WHEN enabled THEN $2 ELSE 'paused' END,last_release_id=$3,last_asset_id=$4,last_tag=$5,last_digest=$6,last_error='',updated_at=NOW() WHERE id=$1`, source.ID, status, release.ID, asset.ID, release.Tag, digest)
	if err != nil {
		return
	}
	_ = tx.Commit()
}
func (s *Service) recoverRun(ctx context.Context, source Source) error {
	rows, err := s.Store.DB.QueryContext(ctx, `SELECT id::text,COALESCE(batch_id::text,''),release_id,asset_id,tag,digest FROM project_doc_source_runs WHERE source_id=$1 AND status='running' ORDER BY started_at`, source.ID)
	if err != nil {
		return err
	}
	type prior struct {
		id, batch      string
		release, asset int64
		tag, digest    string
	}
	var priorRuns []prior
	for rows.Next() {
		var p prior
		if err = rows.Scan(&p.id, &p.batch, &p.release, &p.asset, &p.tag, &p.digest); err != nil {
			rows.Close()
			return err
		}
		priorRuns = append(priorRuns, p)
	}
	err = rows.Err()
	rows.Close()
	if err != nil {
		return err
	}
	for _, run := range priorRuns {
		if run.batch != "" {
			batch, _, e := s.catalog.GetImport(ctx, run.batch)
			if e == nil && batch.Status == "completed" {
				if e = s.syncScopes(ctx, run.batch, source.CollectionID); e != nil {
					return e
				}
				s.complete(ctx, source, run.id, Release{ID: run.release, Tag: run.tag}, ReleaseAsset{ID: run.asset}, run.digest, "completed")
				continue
			}
		}
		_, err = s.Store.DB.ExecContext(ctx, `UPDATE project_doc_source_runs SET status='failed',error_message='上次任务中断，将重新预检。',completed_at=NOW() WHERE id=$1`, run.id)
		if err != nil {
			return err
		}
	}
	return nil
}
