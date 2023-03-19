-- Promotions
INSERT INTO promotions (code, description, is_active, start_date) VALUES 
    ('SESSION_CONFIRMED', 'Confirm the session started', true, now());

-- Invoice requests
ALTER TABLE invoice_requests 
    ADD COLUMN session_id BIGINT;

ALTER TABLE invoice_requests 
    ADD CONSTRAINT fk_invoice_requests_session_id
    FOREIGN KEY (session_id) 
    REFERENCES sessions(id) 
    ON DELETE CASCADE;

-- Sessions
ALTER TABLE sessions 
    ADD COLUMN is_confirmed BOOLEAN NOT NULL DEFAULT false;
