-- Channel Requests
CREATE TYPE channel_request_status AS ENUM (
    'REQUESTED', 
    'AWAITING_PAYMENTS',
    'SETTLING_HTLCS',
    'OPENING_CHANNEL',
    'COMPLETED',
    'FAILED'
);

CREATE TABLE IF NOT EXISTS channel_requests (
    id            BIGSERIAL PRIMARY KEY,
    status        channel_request_status NOT NULL,
    pubkey        BYTEA NOT NULL,
    preimage      BYTEA NOT NULL,
    payment_hash  BYTEA NOT NULL,
    payment_addr  BYTEA NOT NULL,
    amount_msat   BIGINT NOT NULL, 
    settled_msat  BIGINT NOT NULL, 
    funding_tx_id BYTEA,
    output_index  BIGINT
);

ALTER TABLE channel_requests 
    ADD CONSTRAINT uq_channel_requests_payment_hash UNIQUE (payment_hash);

-- Channel Request HTLCs
CREATE TABLE IF NOT EXISTS channel_request_htlcs (
    id                 BIGSERIAL PRIMARY KEY,
    channel_request_id BIGINT NOT NULL,
    chan_id            BIGINT NOT NULL,
    htlc_id            BIGINT NOT NULL,
    is_settled         BOOLEAN NOT NULL
);

ALTER TABLE channel_request_htlcs 
    ADD CONSTRAINT fk_channel_request_htlcs_channel_request_id
    FOREIGN KEY (channel_request_id) 
    REFERENCES channel_requests(id) 
    ON DELETE CASCADE;