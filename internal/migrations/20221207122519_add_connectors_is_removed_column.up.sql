-- Locations
ALTER TABLE locations 
    RENAME COLUMN publish TO is_published;

-- Connectors
ALTER TABLE connectors 
    RENAME COLUMN publish TO is_published;

ALTER TABLE connectors
    ADD COLUMN is_removed BOOLEAN NOT NULL DEFAULT false;

UPDATE connectors SET is_removed = not is_published;
