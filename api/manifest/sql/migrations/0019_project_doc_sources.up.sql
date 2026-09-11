CREATE TABLE project_doc_sources (
    id UUID PRIMARY KEY,
    collection_id UUID NOT NULL UNIQUE REFERENCES collections(id) ON DELETE CASCADE,
    repository TEXT NOT NULL,
    asset_name TEXT NOT NULL DEFAULT 'docs.zip',
    default_locale TEXT NOT NULL,
    mode TEXT NOT NULL DEFAULT 'upsert' CHECK (mode IN ('upsert','replace-version')),
    owner_sub TEXT NOT NULL,
    token_ciphertext TEXT NOT NULL,
    enabled BOOLEAN NOT NULL DEFAULT TRUE,
    status TEXT NOT NULL DEFAULT 'queued',
    last_release_id BIGINT NOT NULL DEFAULT 0,
    last_asset_id BIGINT NOT NULL DEFAULT 0,
    last_tag TEXT NOT NULL DEFAULT '',
    last_digest TEXT NOT NULL DEFAULT '',
    last_error TEXT NOT NULL DEFAULT '',
    checked_at TIMESTAMPTZ,
    next_check_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX ix_project_doc_sources_due ON project_doc_sources(next_check_at) WHERE enabled;
CREATE TABLE project_doc_source_runs (
    id UUID PRIMARY KEY,
    source_id UUID NOT NULL REFERENCES project_doc_sources(id) ON DELETE CASCADE,
    status TEXT NOT NULL,
    stage TEXT NOT NULL DEFAULT 'checking',
    release_id BIGINT NOT NULL DEFAULT 0,
    asset_id BIGINT NOT NULL DEFAULT 0,
    tag TEXT NOT NULL DEFAULT '',
    digest TEXT NOT NULL DEFAULT '',
    batch_id UUID REFERENCES doc_import_batches(id) ON DELETE SET NULL,
    error_message TEXT NOT NULL DEFAULT '',
    started_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    completed_at TIMESTAMPTZ
);
CREATE INDEX ix_project_doc_source_runs_source ON project_doc_source_runs(source_id, started_at DESC);
