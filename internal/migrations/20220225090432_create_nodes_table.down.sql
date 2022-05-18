-- Users
ALTER TABLE IF EXISTS users 
    DROP CONSTRAINT IF EXISTS fk_users_node_id;

ALTER TABLE users
    DROP IF EXISTS node_id;

-- Channel Requests
ALTER TABLE IF EXISTS channel_requests 
    DROP CONSTRAINT IF EXISTS fk_channel_requests_node_id;

ALTER TABLE channel_requests
    DROP IF EXISTS node_id;

-- Nodes
ALTER TABLE IF EXISTS nodes 
  DROP CONSTRAINT IF EXISTS uq_nodes_pubkey;

DROP TABLE IF EXISTS nodes;
