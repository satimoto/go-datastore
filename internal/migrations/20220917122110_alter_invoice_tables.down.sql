-- Sessions
ALTER TABLE IF EXISTS sessions 
    DROP CONSTRAINT IF EXISTS fk_sessions_invoice_request_id;

ALTER TABLE sessions
    DROP IF EXISTS invoice_request_id;

-- Invoice requests
ALTER TABLE invoice_requests
    DROP IF EXISTS tax_msat;

ALTER TABLE invoice_requests
    DROP IF EXISTS tax_fiat;

ALTER TABLE invoice_requests
    DROP IF EXISTS commission_msat;

ALTER TABLE invoice_requests
    DROP IF EXISTS commission_fiat;

ALTER TABLE invoice_requests
    DROP IF EXISTS price_msat;

ALTER TABLE invoice_requests
    DROP IF EXISTS price_fiat;

ALTER TABLE invoice_requests
    DROP IF EXISTS total_fiat;

ALTER TABLE invoice_requests
    DROP IF EXISTS memo;

ALTER TABLE invoice_requests 
    RENAME COLUMN total_msat TO amount_msat;  

-- Session invoices
ALTER TABLE session_invoices
    DROP IF EXISTS total_msat;

ALTER TABLE session_invoices
    DROP IF EXISTS total_fiat;

ALTER TABLE session_invoices 
    RENAME COLUMN price_msat TO amount_msat;  

ALTER TABLE session_invoices 
    RENAME COLUMN price_fiat TO amount_fiat;  
