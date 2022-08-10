-- Users
ALTER TABLE IF EXISTS users 
    DROP CONSTRAINT IF EXISTS fk_users_referrer_id;

ALTER TABLE IF EXISTS users 
    DROP IF EXISTS is_restricted;

ALTER TABLE IF EXISTS users 
    DROP IF EXISTS is_admin;

ALTER TABLE IF EXISTS users 
    DROP IF EXISTS referrer_id;

ALTER TABLE IF EXISTS users 
    DROP IF EXISTS commission_percent;

-- Country accounts
ALTER TABLE IF EXISTS country_accounts 
    DROP CONSTRAINT IF EXISTS uq_country_accounts_country;

DROP INDEX IF EXISTS idx_country_accounts_country;

DROP TABLE IF EXISTS country_accounts;
