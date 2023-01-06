-- Sessions
ALTER TABLE IF EXISTS sessions 
    DROP CONSTRAINT IF EXISTS fk_sessions_invoice_request_id;

ALTER TABLE sessions
    ADD CONSTRAINT fk_sessions_invoice_request_id
    FOREIGN KEY (invoice_request_id) 
    REFERENCES sessions(id) 
    ON DELETE RESTRICT;
