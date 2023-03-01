ALTER TABLE email_subscriptions
    ALTER COLUMN created_date TYPE TIMESTAMPTZ;

ALTER TABLE session_invoices
    ALTER COLUMN last_updated TYPE TIMESTAMPTZ;

ALTER TABLE routing_events
    ALTER COLUMN last_updated TYPE TIMESTAMPTZ;

