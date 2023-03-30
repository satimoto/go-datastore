-- Users
ALTER TABLE users ADD CONSTRAINT uq_users_device_token UNIQUE (device_token);
