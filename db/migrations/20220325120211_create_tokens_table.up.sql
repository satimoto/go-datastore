-- Tokens
CREATE TYPE token_allowed_type AS ENUM (
    'ALLOWED', 
    'BLOCKED', 
    'EXPIRED', 
    'NO_CREDIT', 
    'NOT_ALLOWED'
);

CREATE TYPE token_type AS ENUM (
    'OTHER', 
    'RFID'
);

CREATE TYPE token_whitelist_type AS ENUM (
    'ALWAYS', 
    'ALLOWED', 
    'ALLOWED_OFFLINE', 
    'NEVER'
);

CREATE TABLE IF NOT EXISTS tokens (
    id              BIGSERIAL PRIMARY KEY,
    uid             TEXT NOT NULL,
    type            token_type NOT NULL,
    auth_id         TEXT NOT NULL,
    visual_number   TEXT,
    issuer          TEXT NOT NULL,
    allowed         token_allowed_type NOT NULL,
    valid           BOOLEAN NOT NULL,
    whitelist       token_whitelist_type NOT NULL,
    language        TEXT,
    last_updated    TIMESTAMP NOT NULL
);

CREATE INDEX idx_tokens_uid ON tokens (uid);

ALTER TABLE tokens 
    ADD CONSTRAINT fk_tokens_user_id
    FOREIGN KEY (user_id) 
    REFERENCES users(id) 
    ON DELETE RESTRICT;

-- Users
ALTER TABLE users
    ADD token_id BIGINT;

ALTER TABLE users 
    ADD CONSTRAINT fk_users_token_id
    FOREIGN KEY (token_id) 
    REFERENCES tokens(id)
    ON DELETE SET NULL;