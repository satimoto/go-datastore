-- Price components
ALTER TABLE price_components
    ADD COLUMN currency TEXT;
    
UPDATE price_components SET currency = 'EUR';   

ALTER TABLE price_components
    ALTER COLUMN currency SET NOT NULL;
    