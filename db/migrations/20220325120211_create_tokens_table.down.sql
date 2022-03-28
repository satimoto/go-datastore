-- Users
ALTER TABLE IF EXISTS users 
    DROP CONSTRAINT IF EXISTS fk_users_token_id;

ALTER TABLE users
    DROP IF EXISTS token_id;

-- Tokens
DROP INDEX IF EXISTS idx_tokens_uid;

DROP TABLE IF EXISTS tokens;

DROP TYPE IF EXISTS token_whitelist_type;

DROP TYPE IF EXISTS token_type;

DROP TYPE IF EXISTS token_allowed_type;
