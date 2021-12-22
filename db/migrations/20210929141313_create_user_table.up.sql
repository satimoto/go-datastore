CREATE TYPE authentication_actions AS ENUM (
    'register', 
    'login',
    'link',
    'auth'
);

CREATE TABLE IF NOT EXISTS authentications (
    id BIGSERIAL PRIMARY KEY,
    code TEXT NOT NULL,
    action authentication_actions NOT NULL,
    challenge TEXT NOT NULL,
    signature TEXT,
    linking_key TEXT
);

CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    linking_key TEXT NOT NULL,
    node_key TEXT NOT NULL,
    node_address TEXT NOT NULL,
    device_token TEXT NOT NULL
);

ALTER TABLE users ADD CONSTRAINT uq_users_device_token UNIQUE (device_token);

ALTER TABLE users ADD CONSTRAINT uq_users_node_key UNIQUE (node_key);
