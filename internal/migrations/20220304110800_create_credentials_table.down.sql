ALTER TABLE IF EXISTS credentials 
    DROP CONSTRAINT IF EXISTS uq_credentials_party_country_code;

DROP TABLE IF EXISTS credentials;
