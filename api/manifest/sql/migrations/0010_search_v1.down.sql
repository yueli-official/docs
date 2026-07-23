DROP TABLE IF EXISTS search_batch_receipts;
DROP TABLE IF EXISTS search_documents;
DROP TABLE IF EXISTS search_generations;
DROP TABLE IF EXISTS search_instances;

DROP TRIGGER IF EXISTS docs_search_revision ON docs;
DROP FUNCTION IF EXISTS docs_bump_search_revision();
ALTER TABLE docs DROP COLUMN IF EXISTS search_revision;
ALTER TABLE docs ADD COLUMN search_vector tsvector GENERATED ALWAYS AS (
    setweight(to_tsvector('chinese_zh', coalesce(title, '')), 'A') ||
    setweight(to_tsvector('chinese_zh', coalesce(excerpt, '')), 'B') ||
    setweight(to_tsvector('chinese_zh', coalesce(content, '')), 'C')
) STORED;
CREATE INDEX ix_docs_search ON docs USING GIN (search_vector);
