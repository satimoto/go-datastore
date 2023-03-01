-- Users
ALTER TABLE users
    ADD COLUMN last_active_date TIMESTAMPTZ;
