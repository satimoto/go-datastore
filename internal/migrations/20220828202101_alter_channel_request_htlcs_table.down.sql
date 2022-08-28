-- Channel request htlcs
ALTER TABLE IF EXISTS channel_request_htlcs 
    DROP IF EXISTS is_failed;

ALTER TABLE IF EXISTS channel_request_htlcs 
    DROP IF EXISTS amount_msat;
