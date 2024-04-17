-- Credentials
ALTER TABLE credentials 
    ADD COLUMN is_available BOOLEAN NOT NULL DEFAULT true;
