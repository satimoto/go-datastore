-- Connectors
ALTER TABLE connectors
    ADD COLUMN publish BOOLEAN NOT NULL DEFAULT true;
