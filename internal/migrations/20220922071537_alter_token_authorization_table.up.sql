-- Token authorization
ALTER TABLE token_authorizations
    ADD COLUMN signing_key BYTEA;
