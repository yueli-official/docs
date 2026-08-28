CREATE TABLE collection_locales (
    collection_id UUID NOT NULL REFERENCES collections(id) ON DELETE CASCADE,
    locale        TEXT NOT NULL,
    label         TEXT NOT NULL,
    html_lang     TEXT NOT NULL,
    direction     TEXT NOT NULL DEFAULT 'ltr',
    is_default    BOOLEAN NOT NULL DEFAULT false,
    enabled       BOOLEAN NOT NULL DEFAULT true,
    sort_order    INT NOT NULL DEFAULT 0,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT now(),
    PRIMARY KEY (collection_id, locale),
    CONSTRAINT ck_collection_locales_direction CHECK (direction IN ('ltr', 'rtl')),
    CONSTRAINT ck_collection_locales_locale CHECK (locale ~ '^[A-Za-z]{2,3}(-[A-Za-z0-9]{2,8})*$')
);

CREATE UNIQUE INDEX ux_collection_locales_one_default
    ON collection_locales (collection_id)
    WHERE is_default;

INSERT INTO collection_locales (
    collection_id, locale, label, html_lang, direction, is_default, enabled, sort_order
)
SELECT id, 'en', 'English', 'en', 'ltr', true, true, 0
FROM collections
ON CONFLICT (collection_id, locale) DO NOTHING;

INSERT INTO collection_locales (
    collection_id, locale, label, html_lang, direction, is_default, enabled, sort_order
)
SELECT DISTINCT
    docs.collection_id,
    docs.locale,
    CASE docs.locale
        WHEN 'zh-CN' THEN '简体中文'
        WHEN 'zh-TW' THEN '繁體中文'
        WHEN 'ja' THEN '日本語'
        WHEN 'ko' THEN '한국어'
        ELSE docs.locale
    END,
    docs.locale,
    'ltr',
    false,
    true,
    10
FROM docs
WHERE docs.locale <> 'en'
ON CONFLICT (collection_id, locale) DO NOTHING;

ALTER TABLE docs
    ADD CONSTRAINT fk_docs_collection_locale
    FOREIGN KEY (collection_id, locale)
    REFERENCES collection_locales (collection_id, locale)
    ON UPDATE RESTRICT
    ON DELETE RESTRICT;

CREATE UNIQUE INDEX ux_docs_logical_variant
    ON docs (collection_id, version_id, locale, translation_key)
    WHERE deleted_at IS NULL AND translation_key <> '';
