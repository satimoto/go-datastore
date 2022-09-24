-- Locations
ALTER TABLE locations
    ADD COLUMN publish BOOLEAN NOT NULL DEFAULT false;
