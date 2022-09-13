-- Channel requests
DROP INDEX IF EXISTS channel_point_bytes_idx;

DROP INDEX IF EXISTS channel_point_idx;

ALTER TABLE channel_requests
    DROP IF EXISTS funding_tx_id;

ALTER TABLE channel_requests
   RENAME COLUMN funding_tx_id_bytes TO funding_tx_id;
