ALTER TABLE docs
    DROP CONSTRAINT ck_docs_badge_icon,
    DROP CONSTRAINT ck_docs_badge_text_length,
    DROP COLUMN badge_icon,
    DROP COLUMN badge_text;

DROP INDEX ux_collections_release_family_semantic_version;
DROP INDEX ix_collections_release_family;

ALTER TABLE collections
    DROP CONSTRAINT ck_collections_semantic_version,
    DROP CONSTRAINT ck_collections_release_version_pair,
    DROP COLUMN derived_from_collection_id,
    DROP COLUMN semantic_version,
    DROP COLUMN release_family_id;

DROP TABLE collection_release_families;
