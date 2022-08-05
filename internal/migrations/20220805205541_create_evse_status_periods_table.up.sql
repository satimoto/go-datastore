-- Evse status periods
CREATE TABLE IF NOT EXISTS evse_status_periods (
    id         BIGSERIAL PRIMARY KEY,
    evse_id    BIGINT NOT NULL,
    status     evse_status NOT NULL,
    start_date TIMESTAMP NOT NULL,
    end_date   TIMESTAMP NOT NULL
);

ALTER TABLE evse_status_periods 
    ADD CONSTRAINT fk_evse_status_periods_evse_id 
    FOREIGN KEY (evse_id) 
    REFERENCES evses(id) 
    ON DELETE CASCADE;
