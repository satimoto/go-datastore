-- Evses
DROP INDEX IF EXISTS evses_identifier_idx;

-- Connectors
DROP INDEX IF EXISTS connectors_identifier_idx;

-- Channel requests
DROP INDEX IF EXISTS channel_point_bytes_idx;

DROP INDEX IF EXISTS channel_point_idx;

CREATE UNIQUE INDEX channel_point_bytes_idx ON channel_requests (funding_tx_id_bytes);

CREATE UNIQUE INDEX channel_point_idx ON channel_requests (funding_tx_id);

-- Tokens
DROP INDEX IF EXISTS tokens_uid_idx;

CREATE INDEX idx_tokens_uid ON tokens (uid);
