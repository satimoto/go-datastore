-- Locations
ALTER TABLE locations 
    ADD COLUMN is_removed BOOLEAN NOT NULL DEFAULT false;

UPDATE locations SET is_removed = true WHERE total_evses = 0;
