-- Users
ALTER TABLE users DROP CONSTRAINT uq_users_node_pubkey;

ALTER TABLE users ADD node_address TEXT;
ALTER TABLE users ALTER node_address SET DEFAULT '';
ALTER TABLE users ALTER node_address SET NOT NULL;

ALTER TABLE users RENAME linking_pubkey TO linking_key;
ALTER TABLE users RENAME node_pubkey TO node_key;

ALTER TABLE users ADD CONSTRAINT uq_users_node_key UNIQUE (node_key);

-- Authentications
ALTER TABLE authentications RENAME linking_key TO linking_pubkey;
