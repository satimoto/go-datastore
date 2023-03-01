-- Currencies
DROP INDEX IF EXISTS idx_currencies_code;

ALTER TABLE IF EXISTS currencies 
    DROP CONSTRAINT IF EXISTS uq_currencies_code;

DROP TABLE IF EXISTS currencies;
