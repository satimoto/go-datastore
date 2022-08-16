
-- Channel requests
ALTER TABLE channel_requests
    ADD COLUMN pending_chan_id BYTEA;

-- Psbt funding states
CREATE TABLE IF NOT EXISTS psbt_funding_states (
    id          BIGSERIAL PRIMARY KEY,
    node_id     BIGINT NOT NULL,
    base_psbt   BYTEA NOT NULL,    
    psbt        BYTEA NOT NULL,
    funded_psbt BYTEA,
    signed_psbt BYTEA,
    expiry_date TIMESTAMP NOT NULL
);

ALTER TABLE psbt_funding_states 
    ADD CONSTRAINT fk_psbt_funding_states_node_id 
    FOREIGN KEY (node_id) 
    REFERENCES nodes(id) 
    ON DELETE CASCADE;

-- Psbt funding state channel requests
CREATE TABLE IF NOT EXISTS psbt_funding_state_channel_requests (
    psbt_funding_state_id BIGINT NOT NULL,
    channel_request_id    BIGINT NOT NULL
);

ALTER TABLE psbt_funding_state_channel_requests 
    ADD CONSTRAINT fk_psbt_funding_state_channel_requests_psbt_funding_state_id
    FOREIGN KEY (psbt_funding_state_id) 
    REFERENCES psbt_funding_states(id) 
    ON DELETE CASCADE;

ALTER TABLE psbt_funding_state_channel_requests 
    ADD CONSTRAINT fk_psbt_funding_state_channel_requests_channel_request_id
    FOREIGN KEY (channel_request_id) 
    REFERENCES channel_requests(id) 
    ON DELETE CASCADE;
