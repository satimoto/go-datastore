-- Users
ALTER TABLE users 
    ADD COLUMN battery_capacity FLOAT;

ALTER TABLE users 
    ADD COLUMN battery_power_ac FLOAT;

ALTER TABLE users 
    ADD COLUMN battery_power_dc FLOAT;
