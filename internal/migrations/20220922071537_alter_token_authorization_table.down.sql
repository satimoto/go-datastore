-- Token authorization
ALTER TABLE token_authorizations
    DROP IF EXISTS signing_key;
