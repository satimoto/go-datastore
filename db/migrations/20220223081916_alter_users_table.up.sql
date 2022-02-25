-- Authentications
ALTER TABLE authentications 
    RENAME COLUMN linking_key TO linking_pubkey;

-- Users
ALTER TABLE users 
    DROP CONSTRAINT uq_users_node_key;

ALTER TABLE users 
    DROP node_address;

ALTER TABLE users 
    RENAME COLUMN linking_key TO linking_pubkey;
ALTER TABLE users 
    RENAME COLUMN node_key TO pubkey;

ALTER TABLE users 
    ADD CONSTRAINT uq_users_pubkey UNIQUE (pubkey);
