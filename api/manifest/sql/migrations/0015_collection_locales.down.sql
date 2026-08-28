DROP INDEX IF EXISTS ux_docs_logical_variant;
ALTER TABLE docs DROP CONSTRAINT IF EXISTS fk_docs_collection_locale;
DROP INDEX IF EXISTS ux_collection_locales_one_default;
DROP TABLE IF EXISTS collection_locales;
