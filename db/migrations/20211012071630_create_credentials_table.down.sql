ALTER TABLE credentials 
    DROP CONSTRAINT IF EXISTS fk_credentials_business_details_id;

ALTER TABLE credentials 
    DROP CONSTRAINT IF EXISTS uq_credentials_party_country_code;

DROP TABLE IF EXISTS credentials;
