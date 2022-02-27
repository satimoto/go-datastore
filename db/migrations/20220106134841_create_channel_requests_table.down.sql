-- Channel Request HTLCs
ALTER TABLE IF EXISTS channel_request_htlcs 
    DROP CONSTRAINT IF EXISTS fk_channel_request_htlcs_channel_request_id;

DROP TABLE IF EXISTS channel_request_htlcs;

-- Channel Requests
ALTER TABLE IF EXISTS channel_request
    DROP CONSTRAINT IF EXISTS fk_channel_requests_user_id;

ALTER TABLE IF EXISTS channel_requests 
    DROP CONSTRAINT IF EXISTS uq_channel_requests_payment_hash;

DROP TABLE IF EXISTS channel_requests;

DROP TYPE IF EXISTS channel_request_status;

