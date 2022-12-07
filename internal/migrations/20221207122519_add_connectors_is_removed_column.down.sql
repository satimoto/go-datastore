
-- Connectors
UPDATE connectors SET is_published = is_removed;

ALTER TABLE connectors 
    RENAME COLUMN is_published TO publish;

ALTER TABLE connectors
    DROP IF EXISTS is_removed;

-- Locations
ALTER TABLE locations 
    RENAME COLUMN is_published TO publish;
