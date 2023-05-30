
-- Sessions
ALTER TABLE sessions 
    RENAME COLUMN is_confirmed TO is_confirmed_started;  

ALTER TABLE sessions 
    ADD COLUMN is_confirmed_stopped BOOLEAN NOT NULL DEFAULT false;
