-- Promotions
CREATE TABLE IF NOT EXISTS promotions (
    id          BIGSERIAL PRIMARY KEY,
    code        TEXT NOT NULL,
    description TEXT NOT NULL,
    is_active   BOOLEAN NOT NULL,
    start_date  TIMESTAMPTZ, 
    end_date    TIMESTAMPTZ
);

INSERT INTO promotions (code, description, is_active, start_date) VALUES 
    ('CIRCUIT', 'Earn sats with referrals', true, now()), 
    ('REBATE', 'Rebate overpaid invoice', true, now());

-- Invoice Requests
CREATE TABLE IF NOT EXISTS invoice_requests (
    id              BIGSERIAL PRIMARY KEY,
    user_id         BIGINT NOT NULL,
    promotion_id    BIGINT NOT NULL,
    amount_msat     BIGINT NOT NULL, 
    is_settled      BOOLEAN NOT NULL DEFAULT false,
    payment_request TEXT
);

ALTER TABLE invoice_requests 
    ADD CONSTRAINT fk_invoice_requests_user_id
    FOREIGN KEY (user_id) 
    REFERENCES users(id) 
    ON DELETE RESTRICT;

ALTER TABLE invoice_requests 
    ADD CONSTRAINT fk_invoice_requests_promotion_id
    FOREIGN KEY (promotion_id) 
    REFERENCES promotions(id) 
    ON DELETE RESTRICT;

-- Users
ALTER TABLE users 
    DROP CONSTRAINT IF EXISTS fk_users_referrer_id;

ALTER TABLE users 
    RENAME COLUMN referrer_id TO circuit_user_id;  

ALTER TABLE users 
    ADD CONSTRAINT fk_users_circuit_user_id
    FOREIGN KEY (circuit_user_id) 
    REFERENCES users(id) 
    ON DELETE SET NULL;
