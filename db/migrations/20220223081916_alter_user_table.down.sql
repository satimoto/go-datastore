-- Users
ALTER TABLE IF EXISTS users 
    DROP CONSTRAINT IF EXISTS uq_users_node_pubkey;

ALTER TABLE IF EXISTS users 
    ADD node_address TEXT;
ALTER TABLE IF EXISTS users 
    ALTER node_address SET DEFAULT '';
ALTER TABLE IF EXISTS users 
    ALTER node_address SET NOT NULL;

ALTER TABLE IF EXISTS users 
    RENAME COLUMN IF EXISTS linking_pubkey TO linking_key;
ALTER TABLE IF EXISTS users 
    RENAME COLUMN IF EXISTS node_pubkey TO node_key;

ALTER TABLE IF EXISTS users 
    ADD CONSTRAINT IF NOT EXISTS uq_users_node_key UNIQUE (node_key);

-- Authentications
ALTER TABLE IF EXISTS authentications 
    RENAME COLUMN IF EXISTS linking_pubkey TO linking_key;
