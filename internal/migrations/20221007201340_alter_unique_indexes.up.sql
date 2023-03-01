-- Tokens
DROP INDEX idx_tokens_uid;

CREATE UNIQUE INDEX tokens_uid_idx ON tokens (uid);

-- Channel requests
DROP INDEX channel_point_idx;

DROP INDEX channel_point_bytes_idx;

CREATE INDEX channel_point_idx ON channel_requests (funding_tx_id);

CREATE INDEX channel_point_bytes_idx ON channel_requests (funding_tx_id_bytes);

-- Connectors
CREATE INDEX connectors_identifier_idx ON connectors (identifier) WHERE identifier IS NOT NULL;

-- Evses
CREATE INDEX evses_identifier_idx ON evses (identifier) WHERE identifier IS NOT NULL;
