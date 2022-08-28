
-- Channel request htlcs
ALTER TABLE channel_request_htlcs
    ADD COLUMN amount_msat BIGINT NOT NULL DEFAULT 0;

ALTER TABLE channel_request_htlcs
    ADD COLUMN is_failed BOOLEAN NOT NULL DEFAULT false;
