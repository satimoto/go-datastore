-- Users
ALTER TABLE IF EXISTS users 
    DROP CONSTRAINT IF EXISTS fk_users_circuit_user_id;

ALTER TABLE IF EXISTS users 
    RENAME COLUMN circuit_user_id TO referrer_id;  

ALTER TABLE users
    ADD CONSTRAINT fk_users_referrer_id
    FOREIGN KEY (referrer_id) 
    REFERENCES users(id) 
    ON DELETE SET NULL;

-- Invoice Requests
ALTER TABLE IF EXISTS invoice_requests 
    DROP CONSTRAINT IF EXISTS fk_invoice_requests_promotion_id;

ALTER TABLE IF EXISTS invoice_requests 
    DROP CONSTRAINT IF EXISTS fk_invoice_requests_user_id;

DROP TABLE IF EXISTS invoice_requests;

-- Promotions
DROP TABLE IF EXISTS promotions;
