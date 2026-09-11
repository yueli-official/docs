ALTER TABLE collection_locales ADD COLUMN title TEXT NOT NULL DEFAULT '';
UPDATE collection_locales l SET title = c.title FROM collections c WHERE c.id = l.collection_id;
ALTER TABLE collection_locales ALTER COLUMN title DROP DEFAULT;
