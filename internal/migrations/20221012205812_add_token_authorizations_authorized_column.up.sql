-- Token authorizations
ALTER TABLE token_authorizations
    ADD COLUMN authorized BOOLEAN NOT NULL DEFAULT true;
