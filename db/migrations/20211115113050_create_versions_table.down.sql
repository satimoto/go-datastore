-- Version Endpoints
ALTER TABLE version_endpoints 
    DROP CONSTRAINT IF EXISTS fk_version_endpoints_version_id;

DROP TABLE IF EXISTS version_endpoints;

-- Versions
ALTER TABLE versions 
    DROP CONSTRAINT IF EXISTS fk_versions_credentials_id;

DROP TABLE IF EXISTS versions;

