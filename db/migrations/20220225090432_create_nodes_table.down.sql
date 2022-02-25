-- Users
ALTER TABLE IF EXISTS users 
    DROP CONSTRAINT IF EXISTS fk_users_node_id;

ALTER TABLE users
    DROP node_id;

-- Nodes
ALTER TABLE IF EXISTS nodes 
  DROP CONSTRAINT IF EXISTS uq_nodes_pubkey;

DROP TABLE IF EXISTS nodes;
