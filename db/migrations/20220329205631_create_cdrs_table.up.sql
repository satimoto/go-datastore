-- Calibrations
CREATE TABLE IF NOT EXISTS calibrations (
    id                      BIGSERIAL PRIMARY KEY,
    encoding_method         TEXT NOT NULL,
    encoding_method_version INTEGER,
    public_key              TEXT,
    -- calibration_values   []cdr_calibration_values
    url                     TEXT
);

-- Calibration values
CREATE TABLE IF NOT EXISTS calibration_values (
    id             BIGSERIAL PRIMARY KEY,
    calibration_id BIGINT NOT NULL,
    nature         TEXT NOT NULL,
    plain_data     TEXT NOT NULL,
    signed_data    TEXT NOT NULL
);

ALTER TABLE calibration_values
    ADD CONSTRAINT fk_calibration_values_calibration_id
    FOREIGN KEY (calibration_id) 
    REFERENCES calibrations(id) 
    ON DELETE CASCADE;

-- CDRs
CREATE TABLE IF NOT EXISTS cdrs (
    id                  BIGSERIAL PRIMARY KEY,
    uid                 TEXT NOT NULL,
    credential_id       BIGINT NOT NULL,
    country_code        TEXT,
    party_id            TEXT,
    authorization_id    TEXT,
    start_date_time     TIMESTAMP NOT NULL,
    stop_date_time      TIMESTAMP,
    auth_id             TEXT NOT NULL,
    auth_method         auth_method_type NOT NULL,
    token_id            BIGINT NOT NULL,
    location_id         BIGINT NOT NULL,
    evse_id             BIGINT NOT NULL,
    connector_id        BIGINT NOT NULL,
    meter_id            TEXT,
    currency            TEXT NOT NULL,
    -- tariffs          []tariffs
    -- charging_periods []cdr_charging_periods
    calibration_id      BIGINT,
    total_cost          FLOAT NOT NULL,
    total_energy        FLOAT NOT NULL,
    total_time          FLOAT NOT NULL,
    total_parking_time  FLOAT,
    remark              TEXT,
    last_updated        TIMESTAMP NOT NULL
);

CREATE INDEX idx_cdrs_uid ON cdrs (uid);

ALTER TABLE cdrs
    ADD CONSTRAINT fk_cdrs_credential_id
    FOREIGN KEY (credential_id) 
    REFERENCES credentials(id) 
    ON DELETE SET NULL;

ALTER TABLE cdrs
    ADD CONSTRAINT fk_cdrs_token_id
    FOREIGN KEY (token_id) 
    REFERENCES tokens(id) 
    ON DELETE RESTRICT;

ALTER TABLE cdrs
    ADD CONSTRAINT fk_cdrs_location_id
    FOREIGN KEY (location_id) 
    REFERENCES locations(id) 
    ON DELETE RESTRICT;

ALTER TABLE cdrs
    ADD CONSTRAINT fk_cdrs_evse_id
    FOREIGN KEY (evse_id) 
    REFERENCES evses(id) 
    ON DELETE RESTRICT;

ALTER TABLE cdrs
    ADD CONSTRAINT fk_cdrs_connector_id
    FOREIGN KEY (connector_id) 
    REFERENCES connectors(id) 
    ON DELETE RESTRICT;

ALTER TABLE cdrs
    ADD CONSTRAINT fk_cdrs_calibration_id
    FOREIGN KEY (calibration_id) 
    REFERENCES calibrations(id) 
    ON DELETE RESTRICT;

-- Tariffs
ALTER TABLE tariffs
    ADD cdr_id BIGINT;

ALTER TABLE tariffs
    ADD CONSTRAINT fk_tariffs_cdr_id
    FOREIGN KEY (cdr_id) 
    REFERENCES cdrs(id) 
    ON DELETE CASCADE;

-- CDR charging periods
CREATE TABLE IF NOT EXISTS cdr_charging_periods (
    cdr_id             BIGINT NOT NULL,
    charging_period_id BIGINT NOT NULL
);

ALTER TABLE cdr_charging_periods 
    ADD CONSTRAINT fk_cdr_charging_periods_cdr_id
    FOREIGN KEY (cdr_id) 
    REFERENCES cdrs(id) 
    ON DELETE CASCADE;

ALTER TABLE cdr_charging_periods 
    ADD CONSTRAINT fk_cdr_charging_periods_charging_period_id
    FOREIGN KEY (charging_period_id) 
    REFERENCES charging_periods(id) 
    ON DELETE CASCADE;
