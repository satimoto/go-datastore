-- Channel requests
ALTER TABLE IF EXISTS channel_requests 
    DROP IF EXISTS cltv_expiry_delta;
    
ALTER TABLE IF EXISTS channel_requests 
    DROP IF EXISTS fee_proportional_millionths;  

ALTER TABLE IF EXISTS channel_requests 
    DROP IF EXISTS fee_base_msat;  

ALTER TABLE IF EXISTS channel_requests 
    DROP IF EXISTS scid;
