-- Node scids
CREATE TABLE IF NOT EXISTS node_scids (
    id      BIGSERIAL PRIMARY KEY,
    node_id BIGINT NOT NULL,
    scid    BYTEA NOT NULL
);

ALTER TABLE node_scids 
    ADD CONSTRAINT fk_node_scids_node_id
    FOREIGN KEY (node_id) 
    REFERENCES nodes(id) 
    ON DELETE CASCADE;
