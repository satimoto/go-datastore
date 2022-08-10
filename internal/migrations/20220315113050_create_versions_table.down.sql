-- Credentials
ALTER TABLE IF EXISTS credentials
    DROP CONSTRAINT IF EXISTS fk_credentials_version_id;

ALTER TABLE IF EXISTS credentials
    DROP IF EXISTS version_id;

-- Version Endpoints
ALTER TABLE IF EXISTS version_endpoints
    DROP CONSTRAINT IF EXISTS fk_version_endpoints_version_id;

DROP TABLE IF EXISTS version_endpoints;

-- Versions
ALTER TABLE IF EXISTS versions 
    DROP CONSTRAINT IF EXISTS fk_versions_credentials_id;

DROP TABLE IF EXISTS versions;

