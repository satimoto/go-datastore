-- Country accounts
ALTER TABLE country_accounts
    DROP IF EXISTS longitude;

ALTER TABLE country_accounts
    DROP IF EXISTS latitude;

ALTER TABLE country_accounts
    DROP IF EXISTS zoom;

-- Locations
ALTER TABLE locations
    DROP IF EXISTS added_date;
