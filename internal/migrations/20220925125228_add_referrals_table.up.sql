-- Users
ALTER TABLE users
    ADD COLUMN referral_code TEXT;

-- Referrals
CREATE TABLE IF NOT EXISTS referrals (
    id           BIGSERIAL PRIMARY KEY,
    promotion_id BIGINT NOT NULL,
    user_id      BIGINT NOT NULL,
    ip_address   TEXT NOT NULL,
    last_updated TIMESTAMPTZ NOT NULL
);

ALTER TABLE referrals 
    ADD CONSTRAINT fk_referrals_promotion_id
    FOREIGN KEY (promotion_id) 
    REFERENCES promotions(id) 
    ON DELETE CASCADE;

ALTER TABLE referrals 
    ADD CONSTRAINT fk_referrals_user_id
    FOREIGN KEY (user_id) 
    REFERENCES users(id) 
    ON DELETE CASCADE;
