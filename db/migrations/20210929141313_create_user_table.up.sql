CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    device_token TEXT NOT NULL,
    node_id BIGSERIAL NOT NULL
);

CREATE TABLE IF NOT EXISTS nodes (
    id BIGSERIAL PRIMARY KEY,
    pubkey TEXT NOT NULL,
    address TEXT NOT NULL
);

ALTER TABLE users ADD CONSTRAINT uq_users_device_token UNIQUE (device_token);

ALTER TABLE users ADD CONSTRAINT fk_users_node_id FOREIGN KEY (node_id) REFERENCES nodes(id) ON DELETE CASCADE;

ALTER TABLE nodes ADD CONSTRAINT uq_nodes_pubkey UNIQUE (pubkey);
