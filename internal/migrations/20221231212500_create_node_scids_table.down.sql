-- Node scids
ALTER TABLE IF EXISTS node_scids 
    DROP CONSTRAINT IF EXISTS fk_node_scids_node_id;

DROP TABLE IF EXISTS node_scids;
