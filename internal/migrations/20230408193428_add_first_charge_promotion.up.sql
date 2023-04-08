-- Promotions
INSERT INTO promotions (code, description, is_active, start_date) VALUES 
    ('21_CHARGES', 'User charged 21 times', true, now());

INSERT INTO promotions (code, description, is_active, start_date) VALUES 
    ('FIRST_USER_CHARGE', 'First time user charge', true, now());

INSERT INTO promotions (code, description, is_active, start_date) VALUES 
    ('FIRST_LOCATION_CHARGE', 'First time location charge', true, now());
