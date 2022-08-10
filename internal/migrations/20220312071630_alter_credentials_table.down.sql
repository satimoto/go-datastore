ALTER TABLE IF EXISTS credentials 
    DROP CONSTRAINT IF EXISTS fk_credentials_business_details_id;

ALTER TABLE IF EXISTS credentials 
    DROP IF EXISTS business_details_id;
