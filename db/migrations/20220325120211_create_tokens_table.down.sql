-- Token authorization connectors
ALTER TABLE IF EXISTS token_authorization_connectors 
    DROP CONSTRAINT IF EXISTS fk_token_authorization_connectors_connector_id;

ALTER TABLE IF EXISTS token_authorization_connectors 
    DROP CONSTRAINT IF EXISTS fk_token_authorization_connectors_token_authorization_id;

DROP TABLE IF EXISTS token_authorization_connectors;

-- Token authorization evses
ALTER TABLE IF EXISTS token_authorization_evses 
    DROP CONSTRAINT IF EXISTS fk_token_authorization_evses_evse_id;

ALTER TABLE IF EXISTS token_authorization_evses 
    DROP CONSTRAINT IF EXISTS fk_token_authorization_evses_token_authorization_id;

DROP TABLE IF EXISTS token_authorization_evses;

-- Token authorizations
ALTER TABLE IF EXISTS token_authorizations 
    DROP CONSTRAINT IF EXISTS fk_token_authorizations_token_id;

DROP TABLE IF EXISTS token_authorizations;

-- Tokens
ALTER TABLE IF EXISTS tokens 
    DROP CONSTRAINT IF EXISTS fk_tokens_user_id;

DROP INDEX IF EXISTS idx_tokens_uid;

DROP TABLE IF EXISTS tokens;

DROP TYPE IF EXISTS token_whitelist_type;

DROP TYPE IF EXISTS token_type;

DROP TYPE IF EXISTS token_allowed_type;
