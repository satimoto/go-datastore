CREATE TABLE IF NOT EXISTS email_subscriptions (
    id BIGSERIAL PRIMARY KEY,
    email TEXT NOT NULL,
    code TEXT NOT NULL,
    is_verified BOOLEAN NOT NULL,
    created_date TIMESTAMP NOT NULL
);

ALTER TABLE email_subscriptions ADD CONSTRAINT uq_email_subscriptions_email UNIQUE (email);
