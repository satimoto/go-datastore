-- Token authorizations
ALTER TABLE token_authorizations
    ADD COLUMN signing_key BYTEA;

-- Session invoices
ALTER TABLE session_invoices
    ADD COLUMN signature BYTEA;
