-- Cdrs
ALTER TABLE cdrs
    ADD COLUMN is_flagged BOOLEAN NOT NULL DEFAULT false;

-- Sessions
ALTER TABLE sessions
    ADD COLUMN is_flagged BOOLEAN NOT NULL DEFAULT false;
