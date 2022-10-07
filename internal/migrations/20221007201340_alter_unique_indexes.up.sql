-- Tokens
DROP INDEX idx_tokens_uid;

CREATE UNIQUE INDEX idx_tokens_uid ON tokens (uid);

-- Channel requests
DROP INDEX channel_point_idx;

DROP INDEX channel_point_bytes_idx;

CREATE INDEX channel_point_idx ON channel_requests (funding_tx_id);

CREATE INDEX channel_point_bytes_idx ON channel_requests (funding_tx_id_bytes);
