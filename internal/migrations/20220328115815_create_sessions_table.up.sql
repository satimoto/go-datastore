-- Sessions
CREATE TYPE auth_method_type AS ENUM (
    'AUTH_REQUEST', 
    'WHITELIST'
);

CREATE TYPE session_status_type AS ENUM (
    'ACTIVE', 
    'COMPLETED', 
    'INVALID', 
    'PENDING'
);

CREATE TABLE IF NOT EXISTS sessions (
    id                  BIGSERIAL PRIMARY KEY,
    uid                 TEXT NOT NULL,
    credential_id       BIGINT NOT NULL,
    country_code        TEXT,
    party_id            TEXT,
    authorization_id    TEXT,
    start_datetime      TIMESTAMP NOT NULL,
    end_datetime        TIMESTAMP,
    kwh                 FLOAT NOT NULL,
    auth_id             TEXT NOT NULL,
    auth_method         auth_method_type NOT NULL,
    user_id             BIGINT NOT NULL,
    token_id            BIGINT NOT NULL,
    location_id         BIGINT NOT NULL,
    evse_id             BIGINT NOT NULL,
    connector_id        BIGINT NOT NULL,
    meter_id            TEXT,
    currency            TEXT NOT NULL,
    -- charging_periods []session_charging_periods
    total_cost          FLOAT,
    status              session_status_type NOT NULL,
    last_updated        TIMESTAMP NOT NULL
);

CREATE INDEX idx_sessions_uid ON sessions (uid);

ALTER TABLE sessions
    ADD CONSTRAINT fk_sessions_credential_id
    FOREIGN KEY (credential_id) 
    REFERENCES credentials(id) 
    ON DELETE SET NULL;

ALTER TABLE sessions
    ADD CONSTRAINT fk_sessions_user_id
    FOREIGN KEY (user_id) 
    REFERENCES users(id) 
    ON DELETE RESTRICT;

ALTER TABLE sessions
    ADD CONSTRAINT fk_sessions_token_id
    FOREIGN KEY (token_id) 
    REFERENCES tokens(id) 
    ON DELETE RESTRICT;

ALTER TABLE sessions
    ADD CONSTRAINT fk_sessions_location_id
    FOREIGN KEY (location_id) 
    REFERENCES locations(id) 
    ON DELETE RESTRICT;

ALTER TABLE sessions
    ADD CONSTRAINT fk_sessions_evse_id
    FOREIGN KEY (evse_id) 
    REFERENCES evses(id) 
    ON DELETE RESTRICT;

ALTER TABLE sessions
    ADD CONSTRAINT fk_sessions_connector_id
    FOREIGN KEY (connector_id) 
    REFERENCES connectors(id) 
    ON DELETE RESTRICT;

-- Charging periods
CREATE TABLE IF NOT EXISTS charging_periods (
    id              BIGSERIAL PRIMARY KEY,
    start_date_time TIMESTAMP NOT NULL
    -- dimensions   []charging_period_dimensions
);

-- Charging period dimensions
CREATE TYPE charging_period_dimension_type AS ENUM (
    'ENERGY', 
    'FLAT', 
    'MAX_CURRENT', 
    'MIN_CURRENT', 
    'PARKING_TIME', 
    'TIME'
);

CREATE TABLE IF NOT EXISTS charging_period_dimensions (
    id                 BIGSERIAL PRIMARY KEY,
    charging_period_id BIGINT NOT NULL,
    type               charging_period_dimension_type NOT NULL,
    volume             FLOAT NOT NULL
);

ALTER TABLE charging_period_dimensions 
    ADD CONSTRAINT fk_charging_period_dimensions_charging_period_id 
    FOREIGN KEY (charging_period_id) 
    REFERENCES charging_periods(id) 
    ON DELETE CASCADE;

-- Session charging periods
CREATE TABLE IF NOT EXISTS session_charging_periods (
    session_id         BIGINT NOT NULL,
    charging_period_id BIGINT NOT NULL
);

ALTER TABLE session_charging_periods 
    ADD CONSTRAINT fk_session_charging_periods_session_id
    FOREIGN KEY (session_id) 
    REFERENCES sessions(id) 
    ON DELETE CASCADE;

ALTER TABLE session_charging_periods 
    ADD CONSTRAINT fk_session_charging_periods_charging_period_id
    FOREIGN KEY (charging_period_id) 
    REFERENCES charging_periods(id) 
    ON DELETE CASCADE;
