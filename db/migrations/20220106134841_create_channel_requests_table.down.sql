ALTER TABLE IF EXISTS channel_requests 
  DROP CONSTRAINT IF EXISTS uq_channel_requests_payment_hash;

DROP TABLE IF EXISTS channel_requests;