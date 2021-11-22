ALTER TABLE IF EXISTS email_subscriptions 
    DROP CONSTRAINT IF EXISTS uq_email_subscriptions_email;

DROP TABLE IF EXISTS email_subscriptions;
