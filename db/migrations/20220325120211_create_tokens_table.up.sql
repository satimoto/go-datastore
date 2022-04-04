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

-- Token authorizations
CREATE TABLE IF NOT EXISTS token_authorizations (
    id                BIGSERIAL PRIMARY KEY,
    token_id          BIGINT NOT NULL,
    authorization_id  TEXT NOT NULL,
    country_code      TEXT,
    party_id          TEXT,
    location_id       TEXT
    -- evse_uids      []token_authorization_evses
    -- connector_uids []token_authorization_connectors
);

ALTER TABLE token_authorizations 
    ADD CONSTRAINT fk_token_authorizations_token_id
    FOREIGN KEY (token_id) 
    REFERENCES tokens(id)
    ON DELETE RESTRICT;

-- Token authorization evses
CREATE TABLE IF NOT EXISTS token_authorization_evses (
    token_authorization_id BIGINT NOT NULL,
    evse_uid               TEXT NOT NULL
);

ALTER TABLE token_authorization_evses 
    ADD CONSTRAINT fk_token_authorization_evses_token_authorization_id
    FOREIGN KEY (token_authorization_id) 
    REFERENCES token_authorizations(id)
    ON DELETE CASCADE;

ALTER TABLE token_authorization_evses 
    ADD CONSTRAINT fk_token_authorization_evses_evse_uid
    FOREIGN KEY (evse_uid) 
    REFERENCES evses(uid)
    ON DELETE CASCADE;

-- Token authorization connectors
CREATE TABLE IF NOT EXISTS token_authorization_connectors (
    token_authorization_id BIGINT NOT NULL,
    connector_uid          TEXT NOT NULL
);

ALTER TABLE token_authorization_connectors 
    ADD CONSTRAINT fk_token_authorization_connectors_token_authorization_id
    FOREIGN KEY (token_authorization_id) 
    REFERENCES token_authorizations(id)
    ON DELETE CASCADE;

ALTER TABLE token_authorization_connectors 
    ADD CONSTRAINT fk_token_authorization_connectors_connector_uid
    FOREIGN KEY (connector_uid) 
    REFERENCES connectors(uid)
    ON DELETE CASCADE;

-- Users
ALTER TABLE users
    ADD token_id BIGINT;

ALTER TABLE users 
    ADD CONSTRAINT fk_users_token_id
    FOREIGN KEY (token_id) 
    REFERENCES tokens(id)
    ON DELETE SET NULL;