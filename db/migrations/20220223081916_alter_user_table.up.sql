-- Authentications
ALTER TABLE authentications RENAME linking_key TO linking_pubkey;

-- Users
ALTER TABLE users DROP CONSTRAINT uq_users_node_key;

ALTER TABLE users DROP IF EXISTS node_address;

ALTER TABLE users RENAME linking_key TO linking_pubkey;
ALTER TABLE users RENAME node_key TO node_pubkey;

ALTER TABLE users ADD CONSTRAINT uq_users_node_pubkey UNIQUE (node_pubkey);
