-- Token authorizations
ALTER TABLE token_authorizations
    DROP IF EXISTS verification_key;
