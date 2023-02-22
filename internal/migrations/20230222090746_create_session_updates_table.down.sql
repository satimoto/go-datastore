-- Session updates
ALTER TABLE IF EXISTS session_updates 
    DROP CONSTRAINT IF EXISTS fk_session_updates_user_id;

ALTER TABLE IF EXISTS session_updates 
    DROP CONSTRAINT IF EXISTS fk_session_updates_session_id;

DROP TABLE IF EXISTS session_updates;
