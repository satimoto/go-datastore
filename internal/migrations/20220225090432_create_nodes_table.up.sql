-- Nodes
CREATE TABLE IF NOT EXISTS nodes (
    id          BIGSERIAL PRIMARY KEY,
    pubkey      TEXT NOT NULL,
    node_addr   TEXT NOT NULL,
    lsp_addr    TEXT NOT NULL,
    alias       TEXT NOT NULL,
    color       TEXT NOT NULL,
    commit_hash TEXT NOT NULL,
    version     TEXT NOT NULL,
    channels    BIGINT NOT NULL DEFAULT 0,
    peers       BIGINT NOT NULL DEFAULT 0
);

ALTER TABLE nodes 
    ADD CONSTRAINT uq_nodes_pubkey UNIQUE (pubkey);

-- Channel Requests
ALTER TABLE channel_requests
    ADD COLUMN node_id BIGINT;
ALTER TABLE channel_requests
    ALTER node_id SET NOT NULL;

ALTER TABLE channel_requests 
    ADD CONSTRAINT fk_channel_requests_node_id
    FOREIGN KEY (node_id) 
    REFERENCES nodes(id)
    ON DELETE CASCADE;

-- Users
ALTER TABLE users
    ADD COLUMN node_id BIGINT;

ALTER TABLE users 
    ADD CONSTRAINT fk_users_node_id
    FOREIGN KEY (node_id) 
    REFERENCES nodes(id);