-- Channel requests
ALTER TABLE channel_requests
    ADD COLUMN scid BYTEA;

ALTER TABLE channel_requests
    ADD COLUMN fee_base_msat BIGINT NOT NULL DEFAULT 0;

ALTER TABLE channel_requests
    ADD COLUMN fee_proportional_millionths BIGINT NOT NULL DEFAULT 0;

ALTER TABLE channel_requests
    ADD COLUMN cltv_expiry_delta BIGINT NOT NULL DEFAULT 0;
