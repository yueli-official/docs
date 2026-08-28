package docscomments

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/lib/pq"
)

type Postgres struct{ db *sql.DB }

func NewPostgres(db *sql.DB) *Postgres { return &Postgres{db: db} }

const documentPathsCTE = `WITH RECURSIVE doc_paths AS (
    SELECT id, parent_id, slug, slug::text AS slug_path
    FROM docs
    WHERE parent_id IS NULL
    UNION ALL
    SELECT child.id, child.parent_id, child.slug,
           (parent.slug_path || '/' || child.slug)::text
    FROM docs child
    JOIN doc_paths parent ON parent.id = child.parent_id
)
`

const commentColumns = `id::text, document_id::text,
       COALESCE(parent_id::text, ''), user_sub, author_name,
       content, status, created_at, updated_at`

const qualifiedCommentColumns = `comment.id::text, comment.document_id::text,
       COALESCE(comment.parent_id::text, ''), comment.user_sub, comment.author_name,
       comment.content, comment.status, comment.created_at, comment.updated_at`

func (store *Postgres) PublishedDocument(
	ctx context.Context,
	id string,
) (Document, bool, error) {
	var document Document
	err := store.db.QueryRowContext(ctx, documentPathsCTE+`
SELECT d.id::text, d.title, c.slug, v.key,
       COALESCE(paths.slug_path, d.slug), d.locale
FROM docs d
JOIN collections c ON c.id = d.collection_id
JOIN collection_versions v ON v.id = d.version_id
LEFT JOIN doc_paths paths ON paths.id = d.id
WHERE d.id = $1 AND d.status = 'published' AND v.status = 'published'
  AND d.deleted_at IS NULL`, id).Scan(
		&document.ID, &document.Title, &document.CollectionSlug,
		&document.VersionKey, &document.SlugPath, &document.Locale,
	)
	if err == sql.ErrNoRows {
		return Document{}, false, nil
	}
	return document, err == nil, err
}

func (store *Postgres) Comment(ctx context.Context, id string) (*Comment, error) {
	row := store.db.QueryRowContext(ctx, `SELECT `+commentColumns+`
FROM doc_comments WHERE id = $1 AND deleted_at IS NULL`, id)
	comment, err := scanComment(row.Scan)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return comment, err
}

func (store *Postgres) Insert(ctx context.Context, comment *Comment) error {
	var parent any
	if comment.ParentID != "" {
		parent = comment.ParentID
	}
	_, err := store.db.ExecContext(ctx, `
INSERT INTO doc_comments (
    id, document_id, parent_id, user_sub, author_name, content,
    status, created_at, updated_at
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
		comment.ID, comment.DocumentID, parent, comment.UserSub,
		comment.AuthorName, comment.Content, comment.Status,
		comment.CreatedAt, comment.UpdatedAt,
	)
	return err
}

func (store *Postgres) ApprovedTop(
	ctx context.Context,
	documentID string,
	ascending bool,
	limit int,
	offset int,
) ([]*Comment, int, error) {
	var total int
	if err := store.db.QueryRowContext(ctx, `
SELECT COUNT(*) FROM doc_comments
WHERE document_id = $1 AND parent_id IS NULL
  AND status = 'approved' AND deleted_at IS NULL`, documentID).Scan(&total); err != nil {
		return nil, 0, err
	}
	order := "created_at DESC, id DESC"
	if ascending {
		order = "created_at ASC, id ASC"
	}
	rows, err := store.db.QueryContext(ctx, `SELECT `+commentColumns+`
FROM doc_comments
WHERE document_id = $1 AND parent_id IS NULL
  AND status = 'approved' AND deleted_at IS NULL
ORDER BY `+order+`
LIMIT $2 OFFSET $3`, documentID, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()
	comments, err := scanComments(rows)
	return comments, total, err
}

func (store *Postgres) ApprovedReplies(
	ctx context.Context,
	parentIDs []string,
) (map[string][]*Comment, error) {
	result := map[string][]*Comment{}
	if len(parentIDs) == 0 {
		return result, nil
	}
	rows, err := store.db.QueryContext(ctx, `SELECT `+commentColumns+`
FROM doc_comments
WHERE parent_id = ANY($1::uuid[]) AND status = 'approved' AND deleted_at IS NULL
ORDER BY created_at ASC, id ASC`, pq.Array(parentIDs))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	comments, err := scanComments(rows)
	if err != nil {
		return nil, err
	}
	for _, comment := range comments {
		result[comment.ParentID] = append(result[comment.ParentID], comment)
	}
	return result, nil
}

func (store *Postgres) AdminComments(
	ctx context.Context,
	query AdminQuery,
) ([]*AdminComment, int, error) {
	conditions := []string{"comment.deleted_at IS NULL"}
	arguments := []any{}
	if query.Status != "" {
		arguments = append(arguments, query.Status)
		conditions = append(conditions, fmt.Sprintf("comment.status = $%d", len(arguments)))
	}
	if query.Q != "" {
		arguments = append(arguments, "%"+query.Q+"%")
		placeholder := fmt.Sprintf("$%d", len(arguments))
		conditions = append(conditions, `(
            comment.content ILIKE `+placeholder+` OR
            comment.author_name ILIKE `+placeholder+` OR
            doc.title ILIKE `+placeholder+` OR
            collection.title ILIKE `+placeholder+`
        )`)
	}
	where := strings.Join(conditions, " AND ")
	var total int
	if err := store.db.QueryRowContext(ctx, `
SELECT COUNT(*)
FROM doc_comments comment
JOIN docs doc ON doc.id = comment.document_id
JOIN collections collection ON collection.id = doc.collection_id
WHERE `+where, arguments...).Scan(&total); err != nil {
		return nil, 0, err
	}
	arguments = append(arguments, query.Size, (query.Page-1)*query.Size)
	limitPlaceholder := fmt.Sprintf("$%d", len(arguments)-1)
	offsetPlaceholder := fmt.Sprintf("$%d", len(arguments))
	order := "comment.created_at DESC, comment.id DESC"
	if query.SortOrder == "asc" {
		order = "comment.created_at ASC, comment.id ASC"
	}
	rows, err := store.db.QueryContext(ctx, documentPathsCTE+`
SELECT `+qualifiedCommentColumns+`,
       doc.id::text, doc.title, collection.slug, version.key,
       COALESCE(paths.slug_path, doc.slug), doc.locale
FROM doc_comments comment
JOIN docs doc ON doc.id = comment.document_id
JOIN collections collection ON collection.id = doc.collection_id
JOIN collection_versions version ON version.id = doc.version_id
LEFT JOIN doc_paths paths ON paths.id = doc.id
WHERE `+where+`
ORDER BY `+order+`
LIMIT `+limitPlaceholder+` OFFSET `+offsetPlaceholder, arguments...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()
	items := []*AdminComment{}
	for rows.Next() {
		comment := &Comment{}
		document := Document{}
		if err := rows.Scan(
			&comment.ID, &comment.DocumentID, &comment.ParentID,
			&comment.UserSub, &comment.AuthorName, &comment.Content,
			&comment.Status, &comment.CreatedAt, &comment.UpdatedAt,
			&document.ID, &document.Title, &document.CollectionSlug,
			&document.VersionKey, &document.SlugPath, &document.Locale,
		); err != nil {
			return nil, 0, err
		}
		items = append(items, &AdminComment{Comment: comment, Document: document})
	}
	return items, total, rows.Err()
}

func (store *Postgres) SetStatus(
	ctx context.Context,
	id string,
	status Status,
	updatedAt time.Time,
) (*Comment, error) {
	row := store.db.QueryRowContext(ctx, `
UPDATE doc_comments SET status = $2, updated_at = $3
WHERE id = $1 AND deleted_at IS NULL
RETURNING `+commentColumns, id, status, updatedAt)
	comment, err := scanComment(row.Scan)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return comment, err
}

func (store *Postgres) SoftDelete(
	ctx context.Context,
	id string,
	deletedAt time.Time,
) (bool, error) {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return false, err
	}
	defer tx.Rollback()
	result, err := tx.ExecContext(ctx, `
UPDATE doc_comments SET deleted_at = $2, updated_at = $2
WHERE id = $1 AND deleted_at IS NULL`, id, deletedAt)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil || affected == 0 {
		return false, err
	}
	if _, err := tx.ExecContext(ctx, `
UPDATE doc_comments SET deleted_at = $2, updated_at = $2
WHERE parent_id = $1 AND deleted_at IS NULL`, id, deletedAt); err != nil {
		return false, err
	}
	if err := tx.Commit(); err != nil {
		return false, err
	}
	return true, nil
}

type scanFunc func(...any) error

func scanComment(scan scanFunc) (*Comment, error) {
	comment := &Comment{}
	err := scan(
		&comment.ID, &comment.DocumentID, &comment.ParentID,
		&comment.UserSub, &comment.AuthorName, &comment.Content,
		&comment.Status, &comment.CreatedAt, &comment.UpdatedAt,
	)
	return comment, err
}

func scanComments(rows *sql.Rows) ([]*Comment, error) {
	comments := []*Comment{}
	for rows.Next() {
		comment, err := scanComment(rows.Scan)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, rows.Err()
}
