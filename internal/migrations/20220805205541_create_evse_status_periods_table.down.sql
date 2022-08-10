-- Evse status periods
ALTER TABLE IF EXISTS evse_status_periods 
    DROP CONSTRAINT IF EXISTS fk_evse_status_periods_evse_id;

DROP TABLE IF EXISTS evse_status_periods;
