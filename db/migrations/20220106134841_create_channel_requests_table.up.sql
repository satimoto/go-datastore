
CREATE TABLE IF NOT EXISTS channel_requests (
    id BIGSERIAL PRIMARY KEY,
    pubkey BYTEA NOT NULL,
    payment_hash BYTEA NOT NULL,
    payment_addr BYTEA NOT NULL,
    amount_msat BIGINT NOT NULL, 
    funding_tx_id BYTEA,
    output_index BIGINT
);
