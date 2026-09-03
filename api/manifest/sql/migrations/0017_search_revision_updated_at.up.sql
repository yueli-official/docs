CREATE OR REPLACE FUNCTION docs_bump_search_revision() RETURNS trigger LANGUAGE plpgsql AS $$
BEGIN
    IF (NEW.title, NEW.excerpt, NEW.content, NEW.collection_id, NEW.version_id, NEW.status, NEW.locale, NEW.deleted_at, NEW.updated_at)
       IS DISTINCT FROM
       (OLD.title, OLD.excerpt, OLD.content, OLD.collection_id, OLD.version_id, OLD.status, OLD.locale, OLD.deleted_at, OLD.updated_at) THEN
        NEW.search_revision := OLD.search_revision + 1;
    END IF;
    RETURN NEW;
END$$;
