-- Referrals
ALTER TABLE IF EXISTS referrals 
    DROP CONSTRAINT IF EXISTS fk_referrals_promotion_id;

ALTER TABLE IF EXISTS referrals 
    DROP CONSTRAINT IF EXISTS fk_referrals_user_id;

DROP TABLE IF EXISTS referrals;
