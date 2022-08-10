-- Country accounts
CREATE TABLE IF NOT EXISTS country_accounts (
    id              BIGSERIAL PRIMARY KEY,
    country         TEXT NOT NULL,
    tax_percent     FLOAT NOT NULL
);

CREATE INDEX idx_country_accounts_country ON country_accounts (country);

ALTER TABLE country_accounts 
    ADD CONSTRAINT uq_country_accounts_country UNIQUE (country);

-- Users
ALTER TABLE users 
    ADD COLUMN commission_percent FLOAT NOT NULL DEFAULT 7;
ALTER TABLE users
    ADD COLUMN is_admin BOOLEAN NOT NULL DEFAULT false;
ALTER TABLE users
    ADD COLUMN is_restricted BOOLEAN NOT NULL DEFAULT true;
ALTER TABLE users
    ADD COLUMN referrer_id BIGINT;

ALTER TABLE users
    ADD CONSTRAINT fk_users_referrer_id
    FOREIGN KEY (referrer_id) 
    REFERENCES users(id) 
    ON DELETE SET NULL;
