CREATE TABLE doc_import_batches (
    id             UUID PRIMARY KEY,
    collection_id  UUID NOT NULL REFERENCES collections(id) ON DELETE CASCADE,
    version_id     UUID NOT NULL REFERENCES collection_versions(id) ON DELETE RESTRICT,
    default_locale TEXT NOT NULL,
    mode           TEXT NOT NULL,
    status         TEXT NOT NULL,
    summary_json   JSONB NOT NULL DEFAULT '{}'::jsonb,
    error_message  TEXT NOT NULL DEFAULT '',
    created_by     TEXT NOT NULL,
    created_at     TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at     TIMESTAMPTZ NOT NULL DEFAULT now(),
    completed_at   TIMESTAMPTZ,
    CONSTRAINT ck_doc_import_batches_mode CHECK (mode IN ('create-only', 'upsert', 'replace-version')),
    CONSTRAINT ck_doc_import_batches_status CHECK (status IN ('uploaded', 'checked', 'running', 'completed', 'failed', 'cancelled', 'rolled_back'))
);

CREATE TABLE doc_import_items (
    id                       UUID PRIMARY KEY,
    batch_id                 UUID NOT NULL REFERENCES doc_import_batches(id) ON DELETE CASCADE,
    locale                   TEXT NOT NULL,
    version_key              TEXT NOT NULL,
    path                     TEXT NOT NULL,
    source_markdown_path     TEXT NOT NULL,
    title                    TEXT NOT NULL DEFAULT '',
    slug                     TEXT NOT NULL DEFAULT '',
    translation_key          TEXT NOT NULL DEFAULT '',
    action                   TEXT NOT NULL,
    target_doc_id            UUID REFERENCES docs(id) ON DELETE SET NULL,
    before_doc_json          JSONB NOT NULL DEFAULT '{}'::jsonb,
    after_doc_json           JSONB NOT NULL DEFAULT '{}'::jsonb,
    issues_json              JSONB NOT NULL DEFAULT '[]'::jsonb,
    source_content_hash      TEXT NOT NULL DEFAULT '',
    transformed_content_hash TEXT NOT NULL DEFAULT '',
    created_at               TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT ck_doc_import_items_action CHECK (action IN ('create', 'update', 'skip', 'archive', 'conflict', 'error'))
);

CREATE TABLE doc_import_assets (
    id            UUID PRIMARY KEY,
    batch_id      UUID NOT NULL REFERENCES doc_import_batches(id) ON DELETE CASCADE,
    source_path   TEXT NOT NULL,
    data          BYTEA NOT NULL DEFAULT '\x',
    asset_url     TEXT NOT NULL DEFAULT '',
    content_hash  TEXT NOT NULL DEFAULT '',
    status        TEXT NOT NULL,
    error_message TEXT NOT NULL DEFAULT '',
    created_at    TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT ck_doc_import_assets_status CHECK (status IN ('pending', 'uploaded', 'error'))
);

CREATE TABLE doc_import_asset_refs (
    id                 UUID PRIMARY KEY,
    batch_id           UUID NOT NULL REFERENCES doc_import_batches(id) ON DELETE CASCADE,
    asset_id           UUID REFERENCES doc_import_assets(id) ON DELETE SET NULL,
    item_id            UUID NOT NULL REFERENCES doc_import_items(id) ON DELETE CASCADE,
    markdown_file_path TEXT NOT NULL,
    original_ref       TEXT NOT NULL,
    rewritten_ref      TEXT NOT NULL DEFAULT ''
);

CREATE INDEX ix_doc_import_batches_collection ON doc_import_batches (collection_id, created_at DESC);
CREATE INDEX ix_doc_import_items_batch ON doc_import_items (batch_id, locale, path);
CREATE INDEX ix_doc_import_assets_batch ON doc_import_assets (batch_id, source_path);
CREATE INDEX ix_doc_import_asset_refs_item ON doc_import_asset_refs (item_id);
