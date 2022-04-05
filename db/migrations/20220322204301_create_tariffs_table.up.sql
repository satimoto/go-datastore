-- Tariffs
CREATE TABLE IF NOT EXISTS tariffs (
    id                 BIGSERIAL PRIMARY KEY,
    uid                TEXT NOT NULL,
    country_code       TEXT,
    party_id           TEXT,
    currency           TEXT NOT NULL,
    -- tariff_alt_text []display_texts
    tariff_alt_url     TEXT,
    -- elements        []tariff_elements
    energy_mix_id      BIGINT,
    last_updated       TIMESTAMP NOT NULL
);

CREATE INDEX idx_tariffs_uid ON tariffs (uid);

ALTER TABLE tariffs
    ADD CONSTRAINT fk_tariffs_energy_mix_id
    FOREIGN KEY (energy_mix_id) 
    REFERENCES energy_mixes(id) 
    ON DELETE SET NULL;

-- Tariff Alt Texts
CREATE TABLE IF NOT EXISTS tariff_alt_texts (
    tariff_id       BIGINT NOT NULL,
    display_text_id BIGINT NOT NULL
);

ALTER TABLE tariff_alt_texts 
    ADD CONSTRAINT fk_tariff_alt_texts_tariff_id 
    FOREIGN KEY (tariff_id) 
    REFERENCES tariffs(id) 
    ON DELETE CASCADE;

ALTER TABLE tariff_alt_texts 
    ADD CONSTRAINT fk_tariff_alt_texts_display_text_id
    FOREIGN KEY (display_text_id) 
    REFERENCES display_texts(id) 
    ON DELETE CASCADE;

-- Element Restrictions
CREATE TABLE IF NOT EXISTS element_restrictions (
    id             BIGSERIAL PRIMARY KEY,
    start_time     TEXT,
    end_time       TEXT,
    start_date     TEXT,
    end_date       TEXT,
    min_kwh        FLOAT,
    max_kwh        FLOAT,
    min_power      FLOAT,
    max_power      FLOAT,
    min_duration   INTEGER,
    max_duration   INTEGER
    -- day_of_week []weekdays
);

-- Day Of Week
CREATE TABLE IF NOT EXISTS weekdays (
    id          BIGSERIAL PRIMARY KEY,
    text        TEXT NOT NULL,
    description TEXT NOT NULL
);

INSERT INTO weekdays (text, description) VALUES 
    ('MONDAY', 'Monday'),
    ('TUESDAY', 'Tuesday'),
    ('WEDNESDAY', 'Wednesday'),
    ('THURSDAY', 'Thursday'),
    ('FRIDAY', 'Friday'),
    ('SATURDAY', 'Saturday'),
    ('SUNDAY', 'Sunday');

-- Element Restriction Weekdays
CREATE TABLE IF NOT EXISTS element_restriction_weekdays (
    element_restriction_id BIGINT NOT NULL,
    weekday_id             BIGINT NOT NULL
);

ALTER TABLE element_restriction_weekdays 
    ADD CONSTRAINT fk_element_restriction_weekdays_restriction_id
    FOREIGN KEY (element_restriction_id) 
    REFERENCES element_restrictions(id) 
    ON DELETE CASCADE;

ALTER TABLE element_restriction_weekdays 
    ADD CONSTRAINT fk_element_restriction_weekdays_weekday_id
    FOREIGN KEY (weekday_id) 
    REFERENCES weekdays(id) 
    ON DELETE CASCADE;

-- Tariff Elements
CREATE TABLE IF NOT EXISTS elements (
    id                     BIGSERIAL PRIMARY KEY,
    tariff_id              BIGINT NOT NULL,
    -- price_components    []display_texts
    element_restriction_id BIGINT
);

ALTER TABLE elements 
    ADD CONSTRAINT fk_elements_tariff_id
    FOREIGN KEY (tariff_id) 
    REFERENCES tariffs(id) 
    ON DELETE CASCADE;

ALTER TABLE elements 
    ADD CONSTRAINT fk_elements_element_restriction_id
    FOREIGN KEY (element_restriction_id) 
    REFERENCES element_restrictions(id) 
    ON DELETE SET NULL;

-- Tariff Price Components
CREATE TYPE tariff_dimension AS ENUM (
    'ENERGY', 
    'FLAT',
    'PARKING_TIME',
    'SESSION_TIME',
    'TIME'
);

CREATE TABLE IF NOT EXISTS price_components (
    id         BIGSERIAL PRIMARY KEY,
    element_id BIGINT NOT NULL,
    type       tariff_dimension NOT NULL,
    price      FLOAT NOT NULL,
    step_size  INTEGER NOT NULL
);

ALTER TABLE price_components 
    ADD CONSTRAINT fk_tariff_price_components_element_id
    FOREIGN KEY (element_id) 
    REFERENCES elements(id) 
    ON DELETE CASCADE;
