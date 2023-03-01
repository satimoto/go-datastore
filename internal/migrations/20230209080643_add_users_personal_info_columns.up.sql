-- Users
ALTER TABLE users 
    ADD COLUMN name TEXT;

ALTER TABLE users 
    ADD COLUMN address TEXT;

ALTER TABLE users 
    ADD COLUMN postal_code TEXT;

ALTER TABLE users 
    ADD COLUMN city TEXT;
