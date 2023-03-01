-- Users
ALTER TABLE users
    DROP IF EXISTS city;

ALTER TABLE users
    DROP IF EXISTS address;

ALTER TABLE users
    DROP IF EXISTS postal_code;

ALTER TABLE users
    DROP IF EXISTS name;
