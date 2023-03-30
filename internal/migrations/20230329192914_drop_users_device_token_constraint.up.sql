-- Users
ALTER TABLE IF EXISTS users 
    DROP CONSTRAINT IF EXISTS uq_users_device_token;
