-- Versions
CREATE TABLE IF NOT EXISTS versions (
    id            BIGSERIAL PRIMARY KEY,
    credential_id BIGINT NOT NULL,
    version       TEXT NOT NULL,
    url           TEXT NOT NULL
);

ALTER TABLE versions 
    ADD CONSTRAINT fk_versions_credential_id 
    FOREIGN KEY (credential_id) 
    REFERENCES credentials(id) 
    ON DELETE CASCADE;

-- Version Endpoints
CREATE TABLE IF NOT EXISTS version_endpoints (
    id         BIGSERIAL PRIMARY KEY,
    version_id BIGINT NOT NULL,
    identifier TEXT NOT NULL,
    url        TEXT NOT NULL
);

ALTER TABLE version_endpoints 
    ADD CONSTRAINT fkversion_endpoints_version_id 
    FOREIGN KEY (version_id) 
    REFERENCES versions(id) 
    ON DELETE CASCADE;

-- Credentials
ALTER TABLE credentials
    ADD version_id BIGINT;

ALTER TABLE credentials 
    ADD CONSTRAINT fk_credentials_version_id
    FOREIGN KEY (version_id) 
    REFERENCES versions(id)
    ON DELETE SET NULL;