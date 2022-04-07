-- EVSEs
CREATE TYPE evse_status AS ENUM (
    'AVAILABLE', 
    'BLOCKED',
    'CHARGING',
    'INOPERATIVE',
    'OUTOFORDER',
    'PLANNED',
    'REMOVED',
    'RESERVED',
    'UNKNOWN'
);

CREATE TABLE IF NOT EXISTS capabilities (
    id BIGSERIAL PRIMARY KEY,
    text TEXT NOT NULL,
    description TEXT NOT NULL
);

INSERT INTO capabilities (text, description) VALUES 
    ('CHARGING_PROFILE_CAPABLE', 'Charging Profiles'),
    ('CREDIT_CARD_PAYABLE', 'Credit Cards Accepted'),
    ('REMOTE_START_STOP_CAPABLE', 'Remote Start/Stop'),
    ('RESERVABLE', 'Reservable'),
    ('RFID_READER', 'RFID Tokens Accepted'),
    ('UNLOCK_CAPABLE', 'Remove Unlock');

CREATE TABLE IF NOT EXISTS parking_restrictions (
    id          BIGSERIAL PRIMARY KEY,
    text        TEXT NOT NULL,
    description TEXT NOT NULL
);

INSERT INTO parking_restrictions (text, description) VALUES 
    ('EV_ONLY', 'Reserved parking spot for electric vehicles'),
    ('PLUGGED', 'Parking is only allowed while plugged in'),
    ('DISABLED', 'Reserved parking spot for disabled people'),
    ('CUSTOMERS', 'Parking spot for customers/guests only'),
    ('RFID_READER', 'RFID Tokens Accepted'),
    ('MOTORCYCLES', 'Parking spot for motorcycles or scooters');

CREATE TABLE IF NOT EXISTS evses (
    id                      BIGSERIAL PRIMARY KEY,
    location_id             BIGINT NOT NULL,
    uid                     TEXT NOT NULL,
    evse_id                 TEXT,
    status                  evse_status NOT NULL,
    -- status_schedule      []status_schedules
    -- capabilities         []evse_capabilities 
    -- connectors           []connectors
    floor_level             TEXT,
    geom                    GEOMETRY(POINT, 4326),
    geo_location_id         BIGINT,
    is_remote_capable       BOOLEAN NOT NULL,
    is_rfid_capable         BOOLEAN NOT NULL,
    physical_reference      TEXT,
    -- directions           []display_texts
    -- parking_restrictions []evse_parking_restrictions
    -- images               []images
    last_updated            TIMESTAMP NOT NULL
);

CREATE INDEX idx_evses_uid ON evses (uid);

ALTER TABLE evses 
    ADD CONSTRAINT fk_evses_location_id 
    FOREIGN KEY (location_id) 
    REFERENCES locations(id) 
    ON DELETE CASCADE;

-- Connectors
CREATE TYPE connector_type AS ENUM (
    'CHADEMO', 
    'DOMESTIC_A',
    'DOMESTIC_B',
    'DOMESTIC_C',
    'DOMESTIC_D',
    'DOMESTIC_E',
    'DOMESTIC_F',
    'DOMESTIC_G',
    'DOMESTIC_H',
    'DOMESTIC_I',
    'DOMESTIC_J',
    'DOMESTIC_K',
    'DOMESTIC_L',
    'IEC_60309_2_single_16',
    'IEC_60309_2_three_16',
    'IEC_60309_2_three_32',
    'IEC_60309_2_three_64',
    'IEC_62196_T1',
    'IEC_62196_T1_COMBO',
    'IEC_62196_T2',
    'IEC_62196_T2_COMBO',
    'IEC_62196_T3A',
    'IEC_62196_T3C',
    'TESLA_R',
    'TESLA_S'
);

CREATE TYPE connector_format AS ENUM (
    'SOCKET',
    'CABLE'
);

CREATE TYPE power_type AS ENUM (
    'AC_1_PHASE',
    'AC_3_PHASE',
    'DC'
);

CREATE TABLE IF NOT EXISTS connectors (
    id                   BIGSERIAL PRIMARY KEY,
    evse_id              BIGINT NOT NULL,
    uid                  TEXT NOT NULL,
    standard             connector_type NOT NULL,
    format               connector_format NOT NULL,
    power_type           power_type NOT NULL,
    voltage              INTEGER NOT NULL,
    amperage             INTEGER NOT NULL,
    tariff_id            TEXT,
    terms_and_conditions TEXT,
    last_updated         TIMESTAMP NOT NULL
);

CREATE INDEX idx_connectors_uid ON connectors (uid);

ALTER TABLE connectors 
    ADD CONSTRAINT uq_connectors_evse_uid UNIQUE (evse_id, uid);

ALTER TABLE connectors 
    ADD CONSTRAINT fk_connectors_evse_id 
    FOREIGN KEY (evse_id) 
    REFERENCES evses(id) 
    ON DELETE CASCADE;

-- Evse Directions
CREATE TABLE IF NOT EXISTS evse_directions (
    evse_id         BIGINT NOT NULL,
    display_text_id BIGINT NOT NULL
);

ALTER TABLE evse_directions 
    ADD CONSTRAINT fk_evse_directions_evse_id 
    FOREIGN KEY (evse_id) 
    REFERENCES evses(id) 
    ON DELETE CASCADE;

ALTER TABLE evse_directions 
    ADD CONSTRAINT fk_evse_directions_display_text_id
    FOREIGN KEY (display_text_id) 
    REFERENCES display_texts(id) 
    ON DELETE CASCADE;

-- Evse Images
CREATE TABLE IF NOT EXISTS evse_images (
    evse_id  BIGINT NOT NULL,
    image_id BIGINT NOT NULL
);

ALTER TABLE evse_images
    ADD CONSTRAINT fk_evse_images_location_id 
    FOREIGN KEY (evse_id) 
    REFERENCES evses(id) 
    ON DELETE CASCADE;

ALTER TABLE evse_images 
    ADD CONSTRAINT fk_evse_images_image_id
    FOREIGN KEY (image_id) 
    REFERENCES images(id) 
    ON DELETE CASCADE;

-- Evse Capabilities
CREATE TABLE IF NOT EXISTS evse_capabilities (
    evse_id       BIGINT NOT NULL,
    capability_id BIGINT NOT NULL
);

ALTER TABLE evse_capabilities 
    ADD CONSTRAINT fk_evse_capabilities_evse_id
    FOREIGN KEY (evse_id) 
    REFERENCES evses(id) 
    ON DELETE CASCADE;

ALTER TABLE evse_capabilities 
    ADD CONSTRAINT fk_evse_capabilities_capability_id
    FOREIGN KEY (capability_id) 
    REFERENCES capabilities(id) 
    ON DELETE CASCADE;

-- Evse Parking Restrictions
CREATE TABLE IF NOT EXISTS evse_parking_restrictions (
    evse_id                BIGINT NOT NULL,
    parking_restriction_id BIGINT NOT NULL
);

ALTER TABLE evse_parking_restrictions 
    ADD CONSTRAINT fk_evse_parking_restrictions_evse_id
    FOREIGN KEY (evse_id) 
    REFERENCES evses(id) 
    ON DELETE CASCADE;

ALTER TABLE evse_parking_restrictions 
    ADD CONSTRAINT fk_evse_parking_restrictions_parking_restriction_id
    FOREIGN KEY (parking_restriction_id) 
    REFERENCES parking_restrictions(id) 
    ON DELETE CASCADE;

-- Status Schedules
CREATE TABLE IF NOT EXISTS status_schedules (
    id           BIGSERIAL PRIMARY KEY,
    evse_id      BIGINT NOT NULL,
    period_begin TIMESTAMP NOT NULL,
    period_end   TIMESTAMP,
    status       evse_status NOT NULL
);

ALTER TABLE status_schedules 
    ADD CONSTRAINT fk_status_schedules_evse_id
    FOREIGN KEY (evse_id) 
    REFERENCES evses(id) 
    ON DELETE CASCADE;
