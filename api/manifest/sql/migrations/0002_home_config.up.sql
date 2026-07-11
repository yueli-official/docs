CREATE TABLE home_config (
    key                  TEXT PRIMARY KEY DEFAULT 'default',
    quick_links          JSONB NOT NULL DEFAULT '[]'::jsonb,
    featured_collections JSONB NOT NULL DEFAULT '[]'::jsonb,
    updated_at           TIMESTAMPTZ NOT NULL DEFAULT now(),
    CHECK (key = 'default')
);
