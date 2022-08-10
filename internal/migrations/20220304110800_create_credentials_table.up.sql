-- Credentials
CREATE TABLE IF NOT EXISTS credentials (
    id                 BIGSERIAL PRIMARY KEY,
    client_token       TEXT,
    server_token       TEXT,
    url                TEXT NOT NULL,
    country_code       TEXT NOT NULL,
    party_id           TEXT NOT NULL,
    is_hub             BOOLEAN NOT NULL,
    last_updated       TIMESTAMP NOT NULL
);

ALTER TABLE credentials 
    ADD CONSTRAINT uq_credentials_party_country_code 
    UNIQUE (party_id, country_code);
