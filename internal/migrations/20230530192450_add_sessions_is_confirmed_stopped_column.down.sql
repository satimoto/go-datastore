-- Sessions
ALTER TABLE sessions 
    DROP IF EXISTS is_confirmed_stopped;

ALTER TABLE sessions 
    RENAME COLUMN is_confirmed_started TO is_confirmed;  
