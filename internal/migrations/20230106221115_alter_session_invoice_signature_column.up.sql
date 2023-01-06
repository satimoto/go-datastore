-- Session invoices
ALTER TABLE session_invoices 
    DROP IF EXISTS signature;

ALTER TABLE session_invoices
    ADD COLUMN signature TEXT NOT NULL DEFAULT '';

-- Token authorizations
ALTER TABLE token_authorizations
    DROP IF EXISTS signing_key;
