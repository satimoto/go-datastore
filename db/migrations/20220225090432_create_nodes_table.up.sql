-- Nodes
CREATE TABLE IF NOT EXISTS nodes (
    id            BIGSERIAL PRIMARY KEY,
    pubkey        TEXT NOT NULL,
    alias         TEXT NOT NULL,
    color         TEXT NOT NULL,
    commit_hash   TEXT NOT NULL,
    version       TEXT NOT NULL,
    channels      BIGINT NOT NULL DEFAULT 0,
    peers         BIGINT NOT NULL DEFAULT 0
);

ALTER TABLE nodes 
    ADD CONSTRAINT uq_nodes_pubkey UNIQUE (pubkey);

-- Users
ALTER TABLE users
    ADD node_id BIGINT;

ALTER TABLE users 
    ADD CONSTRAINT fk_users_node_id
    FOREIGN KEY (node_id) 
    REFERENCES nodes(id);