-- Locations
ALTER TABLE locations
    ADD COLUMN added_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP;

-- Country accounts
ALTER TABLE country_accounts
    ADD COLUMN longitude FLOAT;

ALTER TABLE country_accounts
    ADD COLUMN latitude FLOAT;

ALTER TABLE country_accounts
    ADD COLUMN zoom FLOAT;
