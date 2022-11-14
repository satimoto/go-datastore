-- Pending notifications
ALTER TABLE pending_notifications ALTER COLUMN device_token SET NOT NULL;

-- Users
ALTER TABLE users ALTER COLUMN device_token SET NOT NULL;