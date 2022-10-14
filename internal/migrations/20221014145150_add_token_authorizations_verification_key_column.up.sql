-- Token authorizations
ALTER TABLE token_authorizations
    ADD COLUMN verification_key BYTEA;
