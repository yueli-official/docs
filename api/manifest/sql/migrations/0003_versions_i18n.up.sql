CREATE TABLE collection_versions (
    id                UUID PRIMARY KEY,
    collection_id     UUID NOT NULL REFERENCES collections(id) ON DELETE CASCADE,
    key               TEXT NOT NULL,
    label             TEXT NOT NULL,
    status            TEXT NOT NULL DEFAULT 'published',
    is_default        BOOLEAN NOT NULL DEFAULT false,
    sort_order        INT NOT NULL DEFAULT 0,
    source_version_id UUID REFERENCES collection_versions(id) ON DELETE SET NULL,
    created_at        TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at        TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT ck_collection_versions_status CHECK (status IN ('draft', 'published', 'archived')),
    CONSTRAINT ck_collection_versions_key CHECK (key ~ '^(v[0-9][0-9a-z.-]*|next|latest|default)$'),
    UNIQUE (collection_id, key)
);

CREATE UNIQUE INDEX ux_collection_versions_one_default
    ON collection_versions (collection_id)
    WHERE is_default;

ALTER TABLE docs ADD COLUMN version_id UUID REFERENCES collection_versions(id) ON DELETE RESTRICT;
ALTER TABLE docs ADD COLUMN translation_key TEXT NOT NULL;
ALTER TABLE doc_search_events ADD COLUMN version_key TEXT NOT NULL DEFAULT '';

ALTER TABLE docs ALTER COLUMN version_id SET NOT NULL;

ALTER TABLE docs DROP CONSTRAINT docs_collection_id_locale_parent_id_slug_key;
ALTER TABLE docs ADD CONSTRAINT docs_collection_version_locale_parent_slug_key
    UNIQUE (collection_id, version_id, locale, parent_id, slug);

CREATE INDEX ix_docs_collection_version_locale ON docs (collection_id, version_id, locale);
CREATE INDEX ix_docs_translation_key ON docs (translation_key);
CREATE INDEX ix_collection_versions_collection ON collection_versions (collection_id, sort_order, created_at);
