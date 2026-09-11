-- docs site core: tutorial sets (collections) + multi-level doc tree.
CREATE EXTENSION IF NOT EXISTS zhparser;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_ts_config WHERE cfgname = 'chinese_zh') THEN
        CREATE TEXT SEARCH CONFIGURATION chinese_zh (PARSER = zhparser);
        ALTER TEXT SEARCH CONFIGURATION chinese_zh ADD MAPPING FOR n,v,a,i,e,l WITH simple;
    END IF;
END$$;

CREATE TABLE collections (
    id          UUID PRIMARY KEY,
    slug        TEXT NOT NULL UNIQUE,
    title       TEXT NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    cover_url   TEXT NOT NULL DEFAULT '',
    icon        TEXT NOT NULL DEFAULT '',
    sort_order  INT  NOT NULL DEFAULT 0,
    author_sub  TEXT NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE docs (
    id            UUID PRIMARY KEY,
    collection_id UUID NOT NULL REFERENCES collections(id) ON DELETE CASCADE,
    parent_id     UUID REFERENCES docs(id) ON DELETE CASCADE,
    slug          TEXT NOT NULL,
    title         TEXT NOT NULL,
    content       TEXT NOT NULL DEFAULT '',
    excerpt       TEXT NOT NULL DEFAULT '',
    seo_title     TEXT NOT NULL DEFAULT '',
    seo_description TEXT NOT NULL DEFAULT '',
    search_vector tsvector GENERATED ALWAYS AS (
        setweight(to_tsvector('chinese_zh', coalesce(title,   '')), 'A') ||
        setweight(to_tsvector('chinese_zh', coalesce(excerpt, '')), 'B') ||
        setweight(to_tsvector('chinese_zh', coalesce(content, '')), 'C')
    ) STORED,
    status        TEXT NOT NULL DEFAULT 'draft',
    locale        TEXT NOT NULL DEFAULT 'en',
    sort_order    INT  NOT NULL DEFAULT 0,
    author_sub    TEXT NOT NULL,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at    TIMESTAMPTZ,
    UNIQUE (collection_id, locale, parent_id, slug)
);
CREATE INDEX ix_docs_collection_locale ON docs (collection_id, locale);
CREATE INDEX ix_docs_parent ON docs (parent_id);
CREATE INDEX ix_docs_search ON docs USING GIN (search_vector);

CREATE TABLE doc_search_events (
    id              UUID PRIMARY KEY,
    query           TEXT NOT NULL,
    collection_slug TEXT NOT NULL DEFAULT '',
    locale          TEXT NOT NULL DEFAULT 'en',
    result_count    INT  NOT NULL DEFAULT 0,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT now()
);
CREATE INDEX ix_doc_search_events_created ON doc_search_events (created_at DESC);
CREATE INDEX ix_doc_search_events_query ON doc_search_events (query);
