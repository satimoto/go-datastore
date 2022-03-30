-- CDR charging periods
ALTER TABLE IF EXISTS cdr_charging_periods 
    DROP CONSTRAINT IF EXISTS fk_cdr_charging_periods_charging_period_id;

ALTER TABLE IF EXISTS cdr_charging_periods 
    DROP CONSTRAINT IF EXISTS fk_cdr_charging_periods_cdr_id;

DROP TABLE IF EXISTS cdr_charging_periods;

-- Tariffs
ALTER TABLE IF EXISTS tariffs
    DROP CONSTRAINT IF EXISTS fk_tariffs_cdr_id;

ALTER TABLE tariffs
    DROP IF EXISTS cdr_id;

-- CDRs
ALTER TABLE IF EXISTS cdrs 
    DROP CONSTRAINT IF EXISTS fk_cdrs_location_id;

DROP INDEX IF EXISTS idx_cdrs_uid;

DROP TABLE IF EXISTS cdrs;
