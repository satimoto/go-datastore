-- Tariffs
ALTER TABLE tariffs
    ADD COLUMN is_intermediate_cdr_capable BOOLEAN NOT NULL DEFAULT true;

-- Locations
ALTER TABLE locations
    DROP IF EXISTS is_intermediate_cdr_capable;

-- Parties
ALTER TABLE IF EXISTS parties 
    DROP CONSTRAINT IF EXISTS fk_parties_credential_id;

ALTER TABLE IF EXISTS parties 
    DROP CONSTRAINT IF EXISTS uq_parties_country_code_party_id;

DROP TABLE IF EXISTS parties;
