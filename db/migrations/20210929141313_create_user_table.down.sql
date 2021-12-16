
-- Authentications
DROP TABLE IF EXISTS authentications;

DROP TYPE IF EXISTS authentication_actions;

-- Users
ALTER TABLE IF EXISTS users 
  DROP CONSTRAINT IF EXISTS uq_users_node_key;

ALTER TABLE IF EXISTS users 
  DROP CONSTRAINT IF EXISTS uq_users_device_token;

DROP TABLE IF EXISTS users;