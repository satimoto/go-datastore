-- Users
ALTER TABLE IF EXISTS users 
    DROP CONSTRAINT IF EXISTS uq_users_pubkey;

ALTER TABLE IF EXISTS users 
    ADD COLUMN node_address TEXT;

ALTER TABLE IF EXISTS users 
    RENAME COLUMN linking_pubkey TO linking_key;
ALTER TABLE IF EXISTS users 
    RENAME COLUMN pubkey TO node_key;

ALTER TABLE IF EXISTS users 
    ADD CONSTRAINT uq_users_node_key UNIQUE (node_key);

-- Authentications
ALTER TABLE IF EXISTS authentications 
    RENAME COLUMN linking_pubkey TO linking_key;
