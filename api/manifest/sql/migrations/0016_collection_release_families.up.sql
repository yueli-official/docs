CREATE TABLE collection_release_families (
    id          UUID PRIMARY KEY,
    name        TEXT NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

ALTER TABLE collections
    ADD COLUMN release_family_id UUID REFERENCES collection_release_families(id) ON DELETE SET NULL,
    ADD COLUMN semantic_version TEXT,
    ADD COLUMN derived_from_collection_id UUID REFERENCES collections(id) ON DELETE SET NULL;

ALTER TABLE collections
    ADD CONSTRAINT ck_collections_release_version_pair CHECK (
        (release_family_id IS NULL AND semantic_version IS NULL)
        OR (release_family_id IS NOT NULL AND semantic_version IS NOT NULL)
    ),
    ADD CONSTRAINT ck_collections_semantic_version CHECK (
        semantic_version IS NULL
        OR semantic_version ~ '^(0|[1-9][0-9]*)\.(0|[1-9][0-9]*)\.(0|[1-9][0-9]*)$'
    );

CREATE UNIQUE INDEX ux_collections_release_family_semantic_version
    ON collections (release_family_id, semantic_version)
    WHERE release_family_id IS NOT NULL;

CREATE INDEX ix_collections_release_family
    ON collections (release_family_id);

ALTER TABLE docs
    ADD COLUMN badge_text TEXT NOT NULL DEFAULT '',
    ADD COLUMN badge_icon TEXT NOT NULL DEFAULT '';

ALTER TABLE docs
    ADD CONSTRAINT ck_docs_badge_text_length CHECK (char_length(badge_text) <= 24),
    ADD CONSTRAINT ck_docs_badge_icon CHECK (
        badge_icon = '' OR badge_icon ~ '^i-tabler-[a-z0-9]+(-[a-z0-9]+)*$'
    );
