-- Psbt funding state channel requests
ALTER TABLE IF EXISTS psbt_funding_state_channel_requests 
    DROP CONSTRAINT IF EXISTS fk_psbt_funding_state_channel_requests_channel_request_id;

ALTER TABLE IF EXISTS psbt_funding_state_channel_requests 
    DROP CONSTRAINT IF EXISTS fk_psbt_funding_state_channel_requests_psbt_funding_state_id;

DROP TABLE IF EXISTS psbt_funding_state_channel_requests;

-- Psbt funding states
ALTER TABLE IF EXISTS psbt_funding_states 
    DROP CONSTRAINT IF EXISTS fk_psbt_funding_states_node_id;

DROP TABLE IF EXISTS psbt_funding_states;

-- Channel requests
ALTER TABLE IF EXISTS channel_requests 
    DROP IF EXISTS pending_chan_id;
