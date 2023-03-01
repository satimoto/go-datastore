-- Channel requests
ALTER TABLE channel_requests
    RENAME COLUMN funding_tx_id TO funding_tx_id_bytes;

ALTER TABLE channel_requests
    ADD COLUMN funding_tx_id TEXT;

CREATE UNIQUE INDEX channel_point_idx ON channel_requests (funding_tx_id);

CREATE UNIQUE INDEX channel_point_bytes_idx ON channel_requests (funding_tx_id_bytes);
