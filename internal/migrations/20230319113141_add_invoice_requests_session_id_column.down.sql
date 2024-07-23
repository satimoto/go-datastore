-- Sessions
ALTER TABLE sessions 
    DROP IF EXISTS is_confirmed;

-- Invoice requests
ALTER TABLE IF EXISTS invoice_requests 
    DROP CONSTRAINT IF EXISTS fk_invoice_requests_session_id;

ALTER TABLE invoice_requests
    DROP IF EXISTS session_id;

-- Promotions
DELETE FROM promotions
    WHERE code = 'SESSION_CONFIRMED';
