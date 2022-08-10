-- Nodes
ALTER TABLE nodes
    ADD COLUMN is_active boolean NOT NULL DEFAULT true;
