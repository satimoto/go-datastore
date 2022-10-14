-- Sessions
ALTER TABLE sessions
    DROP IF EXISTS is_flagged;

-- Cdrs
ALTER TABLE cdrs
    DROP IF EXISTS is_flagged;
