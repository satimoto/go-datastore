-- CDRs
CREATE TABLE IF NOT EXISTS cdrs (
    id                  BIGSERIAL PRIMARY KEY,
    uid                 TEXT NOT NULL,
    start_date_time     TIMESTAMP NOT NULL,
    stop_date_time      TIMESTAMP,
    auth_id             TEXT NOT NULL,
    auth_method         auth_method_type NOT NULL,
    location_id         BIGINT NOT NULL,
    meter_id            TEXT,
    currency            TEXT NOT NULL,
    -- tariffs          []tariffs
    -- charging_periods []cdr_charging_periods
    total_cost          FLOAT NOT NULL,
    total_energy        FLOAT NOT NULL,
    total_time          FLOAT NOT NULL,
    total_parking_time  FLOAT,
    remark              TEXT,
    last_updated        TIMESTAMP NOT NULL
);

CREATE INDEX idx_cdrs_uid ON cdrs (uid);

ALTER TABLE cdrs
    ADD CONSTRAINT fk_cdrs_location_id
    FOREIGN KEY (location_id) 
    REFERENCES locations(id) 
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
