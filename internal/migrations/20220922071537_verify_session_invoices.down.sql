-- Session invoices
ALTER TABLE session_invoices
    DROP IF EXISTS signature;

-- Token authorizations
ALTER TABLE token_authorizations
    DROP IF EXISTS signing_key;
