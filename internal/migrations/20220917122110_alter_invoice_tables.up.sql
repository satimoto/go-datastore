-- Session invoices
ALTER TABLE session_invoices 
    RENAME COLUMN amount_fiat TO price_fiat;  

ALTER TABLE session_invoices 
    RENAME COLUMN amount_msat TO price_msat;  

ALTER TABLE session_invoices
    ADD COLUMN total_fiat FLOAT NOT NULL;

ALTER TABLE session_invoices
    ADD COLUMN total_msat BIGINT NOT NULL;

-- Invoice requests
ALTER TABLE invoice_requests 
    RENAME COLUMN amount_msat TO total_msat;  

ALTER TABLE invoice_requests
    ADD COLUMN currency TEXT NOT NULL;

ALTER TABLE invoice_requests
    ADD COLUMN memo TEXT NOT NULL;

ALTER TABLE invoice_requests
    ADD COLUMN total_fiat FLOAT NOT NULL;

ALTER TABLE invoice_requests
    ADD COLUMN price_fiat FLOAT;

ALTER TABLE invoice_requests
    ADD COLUMN price_msat BIGINT;

ALTER TABLE invoice_requests
    ADD COLUMN commission_fiat FLOAT;

ALTER TABLE invoice_requests
    ADD COLUMN commission_msat BIGINT;

ALTER TABLE invoice_requests
    ADD COLUMN tax_fiat FLOAT;

ALTER TABLE invoice_requests
    ADD COLUMN tax_msat BIGINT;

-- Sessions
ALTER TABLE sessions
    ADD COLUMN invoice_request_id BIGINT;

ALTER TABLE sessions
    ADD CONSTRAINT fk_sessions_invoice_request_id
    FOREIGN KEY (invoice_request_id) 
    REFERENCES sessions(id) 
    ON DELETE RESTRICT;
