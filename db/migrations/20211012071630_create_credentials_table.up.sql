-- Credentials
CREATE TABLE IF NOT EXISTS credentials (
    id BIGSERIAL PRIMARY KEY,
    client_token TEXT,
    server_token TEXT,
    version TEXT,
    url TEXT NOT NULL,
    business_detail_id BIGINT NOT NULL,
    party_id TEXT NOT NULL,
    country_code TEXT NOT NULL,
    last_updated TIMESTAMP NOT NULL
);

ALTER TABLE credentials 
    ADD CONSTRAINT uq_credentials_party_country_code 
    UNIQUE (party_id, country_code);

ALTER TABLE credentials 
    ADD CONSTRAINT fk_credentials_business_details_id
    FOREIGN KEY (business_detail_id) 
    REFERENCES business_details(id) 
    ON DELETE CASCADE;
