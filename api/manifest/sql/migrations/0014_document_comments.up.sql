CREATE TABLE doc_comments (
    id          UUID PRIMARY KEY,
    document_id UUID NOT NULL REFERENCES docs(id) ON DELETE CASCADE,
    parent_id   UUID REFERENCES doc_comments(id) ON DELETE CASCADE,
    user_sub    TEXT NOT NULL,
    author_name TEXT NOT NULL,
    content     TEXT NOT NULL,
    status      TEXT NOT NULL DEFAULT 'approved'
        CHECK (status IN ('approved', 'pending', 'spam', 'trash')),
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at  TIMESTAMPTZ,
    CHECK (char_length(content) BETWEEN 1 AND 2000)
);

CREATE INDEX ix_doc_comments_document_status_time
    ON doc_comments (document_id, status, created_at DESC)
    WHERE deleted_at IS NULL;

CREATE INDEX ix_doc_comments_parent
    ON doc_comments (parent_id)
    WHERE deleted_at IS NULL;

CREATE INDEX ix_doc_comments_moderation
    ON doc_comments (status, created_at DESC)
    WHERE deleted_at IS NULL;
