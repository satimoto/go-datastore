-- Users
ALTER TABLE users ALTER COLUMN device_token DROP NOT NULL;

-- Pending notifications
ALTER TABLE pending_notifications ALTER COLUMN device_token DROP NOT NULL;