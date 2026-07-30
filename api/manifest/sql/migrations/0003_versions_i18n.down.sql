ALTER TABLE docs DROP CONSTRAINT IF EXISTS docs_collection_version_locale_parent_slug_key;
DROP INDEX IF EXISTS ix_docs_collection_version_locale;
DROP INDEX IF EXISTS ix_docs_translation_key;
DROP INDEX IF EXISTS ix_collection_versions_collection;
DROP INDEX IF EXISTS ux_collection_versions_one_default;

ALTER TABLE doc_search_events DROP COLUMN IF EXISTS version_key;
ALTER TABLE docs DROP COLUMN IF EXISTS translation_key;
ALTER TABLE docs DROP COLUMN IF EXISTS version_id;

ALTER TABLE docs ADD CONSTRAINT docs_collection_id_locale_parent_id_slug_key
    UNIQUE (collection_id, locale, parent_id, slug);

DROP TABLE IF EXISTS collection_versions;
