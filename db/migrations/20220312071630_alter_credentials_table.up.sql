-- Credentials
ALTER TABLE credentials
    ADD business_detail_id BIGINT NOT NULL;

ALTER TABLE credentials 
    ADD CONSTRAINT fk_credentials_business_details_id
    FOREIGN KEY (business_detail_id) 
    REFERENCES business_details(id) 
    ON DELETE CASCADE;
