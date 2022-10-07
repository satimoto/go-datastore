-- Channel requests
DROP INDEX IF EXISTS channel_point_bytes_idx;

DROP INDEX IF EXISTS channel_point_idx;

CREATE UNIQUE INDEX channel_point_bytes_idx ON channel_requests (funding_tx_id_bytes);

CREATE UNIQUE INDEX channel_point_idx ON channel_requests (funding_tx_id);

-- Tokens
DROP INDEX IF EXISTS idx_tokens_uid;

CREATE INDEX idx_tokens_uid ON tokens (uid);
