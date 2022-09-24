-- Locations
ALTER TABLE locations
    ADD COLUMN publish BOOLEAN DEFAULT false;
