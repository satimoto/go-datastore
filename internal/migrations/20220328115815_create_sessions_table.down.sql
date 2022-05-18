-- Session charging periods
ALTER TABLE IF EXISTS session_charging_periods 
    DROP CONSTRAINT IF EXISTS fk_session_charging_periods_charging_period_id;

ALTER TABLE IF EXISTS session_charging_periods 
    DROP CONSTRAINT IF EXISTS fk_session_charging_periods_session_id;

DROP TABLE IF EXISTS session_charging_periods;

-- Charging period dimensions
ALTER TABLE IF EXISTS charging_period_dimensions 
    DROP CONSTRAINT IF EXISTS fk_charging_period_dimensions_charging_period_id;

DROP TABLE IF EXISTS charging_period_dimensions;

DROP TYPE IF EXISTS charging_period_dimension_type;

-- Charging periods
DROP TABLE IF EXISTS charging_periods;

-- Sessions
ALTER TABLE IF EXISTS sessions 
    DROP CONSTRAINT IF EXISTS fk_sessions_connector_id;
    
ALTER TABLE IF EXISTS sessions 
    DROP CONSTRAINT IF EXISTS fk_sessions_evse_id;
    
ALTER TABLE IF EXISTS sessions 
    DROP CONSTRAINT IF EXISTS fk_sessions_location_id;

ALTER TABLE IF EXISTS sessions 
    DROP CONSTRAINT IF EXISTS fk_sessions_token_id;

ALTER TABLE IF EXISTS sessions 
    DROP CONSTRAINT IF EXISTS fk_sessions_user_id;
    
ALTER TABLE IF EXISTS sessions 
    DROP CONSTRAINT IF EXISTS fk_sessions_credential_id;

DROP INDEX IF EXISTS idx_sessions_uid;

DROP TABLE IF EXISTS sessions;

DROP TYPE IF EXISTS session_status_type;

DROP TYPE IF EXISTS auth_method_type;
