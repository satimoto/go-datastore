-- Display Texts
CREATE TABLE IF NOT EXISTS display_texts (
    id BIGSERIAL PRIMARY KEY,
    language TEXT NOT NULL,
    text TEXT NOT NULL
);

-- Geo Locations
CREATE TABLE IF NOT EXISTS geo_locations (
    id BIGSERIAL PRIMARY KEY,
    latitude DECIMAL NOT NULL,
    longitude DECIMAL NOT NULL,
    name TEXT
);

-- Images
CREATE TYPE image_category AS ENUM (
    'CHARGER', 
    'ENTRANCE',
    'LOCATION',
    'NETWORK',
    'OPERATOR',
    'OTHER',
    'OWNER'
);

CREATE TABLE IF NOT EXISTS images (
    id BIGSERIAL PRIMARY KEY,
    url TEXT NOT NULL,
    thumbnail TEXT,
    category image_category NOT NULL,
    type TEXT NOT NULL,
    width SMALLINT,
    height SMALLINT
);

-- Facilities
CREATE TABLE IF NOT EXISTS facilities (
    id BIGSERIAL PRIMARY KEY,
    text TEXT NOT NULL,
    description TEXT NOT NULL
);

INSERT INTO facilities (text, description) VALUES 
    ('AIRPORT', 'Airport'),
    ('BUS_STOP', 'Bus Stop'),
    ('CAFE', 'Cafe'),
    ('CARPOOL_PARKING', 'Carpool Parking'),
    ('FUEL_STATION', 'Fuel Station'),
    ('HOTEL', 'Hotel'),
    ('MALL', 'Mall'),
    ('MUSEUM', 'Museum'),
    ('NATURE', 'Park/Nature Reserve'),
    ('RECREATION_AREA', 'Recreation Area'),
    ('RESTAURANT', 'Restaurant'),
    ('SPORT', 'Sports Facility'),
    ('SUPERMARKET', 'Supermarket'),
    ('TAXI_STAND', 'Taxi Stand'),
    ('TRAIN_STATION', 'Train Station'),
    ('WIFI', 'Wifi');

-- Business Details
CREATE TABLE IF NOT EXISTS business_details (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    website TEXT,
    logo_id BIGSERIAL
);

ALTER TABLE business_details
    ADD CONSTRAINT fk_business_details_logo_id
    FOREIGN KEY (logo_id) 
    REFERENCES images(id) 
    ON DELETE SET NULL;

-- Energy Mix
CREATE TYPE energy_source_category AS ENUM (
    'NUCLEAR', 
    'GENERAL_FOSSIL',
    'COAL',
    'GAS',
    'GENERAL_GREEN',
    'SOLAR',
    'WIND',
    'WATER'
);

CREATE TYPE environmental_impact_category AS ENUM (
    'NUCLEAR_WASTE', 
    'CARBON_DIOXIDE'
);

CREATE TABLE IF NOT EXISTS energy_sources (
    id BIGSERIAL PRIMARY KEY,
    energy_mix_id BIGSERIAL NOT NULL,
    source energy_source_category NOT NULL,
    percentage DECIMAL NOT NULL
);

ALTER TABLE energy_sources 
    ADD CONSTRAINT fk_energy_sources_energy_mix_id 
    FOREIGN KEY (energy_mix_id) 
    REFERENCES energy_mixes(id) 
    ON DELETE CASCADE;

CREATE TABLE IF NOT EXISTS environmental_impacts (
    id BIGSERIAL PRIMARY KEY,
    energy_mix_id BIGSERIAL NOT NULL,
    source environmental_impact_category NOT NULL,
    amount DECIMAL NOT NULL
);

ALTER TABLE environmental_impacts 
    ADD CONSTRAINT fk_environmental_impacts_energy_mix_id 
    FOREIGN KEY (energy_mix_id) 
    REFERENCES energy_mixes(id) 
    ON DELETE CASCADE;

CREATE TABLE IF NOT EXISTS energy_mixes (
    id BIGSERIAL PRIMARY KEY,
    is_green_energy BOOLEAN NOT NULL,
    -- energy_sources []energy_sources
    -- environ_impact []environmental_impacts
    supplier_name TEXT,
    energy_product_name TEXT 
);

-- Location Hours
CREATE TYPE period_type AS ENUM (
    'OPENING', 
    'CLOSING'
);

CREATE TABLE IF NOT EXISTS opening_times (
    id BIGSERIAL PRIMARY KEY,
    -- regular_hours []regular_hours
    twentyfourseven BOOLEAN NOT NULL
    -- exceptional_openings []exceptional_periods
    -- exceptional_closings []exceptional_periods
);

CREATE TABLE IF NOT EXISTS regular_hours (
    id BIGSERIAL PRIMARY KEY,
    opening_time_id BIGSERIAL NOT NULL,
    weekday SMALLINT NOT NULL,
    period_begin TEXT NOT NULL,
    period_end TEXT NOT NULL
);

ALTER TABLE regular_hours 
    ADD CONSTRAINT fk_regular_hours_opening_time_id
    FOREIGN KEY (opening_time_id) 
    REFERENCES opening_times(id) 
    ON DELETE CASCADE;

CREATE TABLE IF NOT EXISTS exceptional_periods (
    id BIGSERIAL PRIMARY KEY,
    opening_time_id BIGSERIAL NOT NULL,
    period_type period_type NOT NULL,
    period_begin TIMESTAMP NOT NULL,
    period_end TIMESTAMP NOT NULL
);

ALTER TABLE exceptional_periods 
    ADD CONSTRAINT fk_exceptional_periods_opening_time_id
    FOREIGN KEY (opening_time_id) 
    REFERENCES opening_times(id) 
    ON DELETE CASCADE;

-- Locations
CREATE TYPE location_type AS ENUM (
    'ON_STREET', 
    'PARKING_GARAGE', 
    'UNDERGROUND_GARAGE', 
    'PARKING_LOT', 
    'OTHER', 
    'UNKNOWN'
);

CREATE TABLE IF NOT EXISTS locations (
    id BIGSERIAL PRIMARY KEY,
    uid TEXT NOT NULL,
    type location_type NOT NULL,
    name TEXT,
    address TEXT NOT NULL,
    city TEXT NOT NULL,
    postal_code TEXT NOT NULL,
    country TEXT NOT NULL,
    geom GEOMETRY(POINT, 4326) NOT NULL,
    geo_location_id BIGSERIAL NOT NULL,
    -- related_locations []related_locations
    -- evses []evses
    -- directions []location_directions
    operator_id BIGSERIAL,
    suboperator_id BIGSERIAL,
    owner_id BIGSERIAL,
    -- facilities []facilities
    time_zone TEXT NOT NULL,
    opening_time_id BIGSERIAL,
    charging_when_closed BOOLEAN NOT NULL,
    -- images []location_images
    energy_mix_id BIGSERIAL,
    last_updated TIMESTAMP NOT NULL
);

CREATE INDEX idx_locations_uid ON locations (uid);

ALTER TABLE locations
    ADD CONSTRAINT fk_locations_operator_id
    FOREIGN KEY (operator_id) 
    REFERENCES business_details(id) 
    ON DELETE SET NULL;

ALTER TABLE locations
    ADD CONSTRAINT fk_locations_suboperator_id
    FOREIGN KEY (suboperator_id) 
    REFERENCES business_details(id) 
    ON DELETE SET NULL;

ALTER TABLE locations
    ADD CONSTRAINT fk_locations_owner_id
    FOREIGN KEY (owner_id) 
    REFERENCES business_details(id) 
    ON DELETE SET NULL;

ALTER TABLE locations
    ADD CONSTRAINT fk_locations_energy_mix_id
    FOREIGN KEY (energy_mix_id) 
    REFERENCES energy_mixes(id) 
    ON DELETE SET NULL;

ALTER TABLE locations
    ADD CONSTRAINT fk_locations_opening_time_id
    FOREIGN KEY (opening_time_id) 
    REFERENCES opening_times(id) 
    ON DELETE SET NULL;

-- Related Locations
CREATE TABLE IF NOT EXISTS related_locations (
    location_id BIGSERIAL NOT NULL,
    geo_location_id BIGSERIAL NOT NULL
);

ALTER TABLE related_locations 
    ADD CONSTRAINT fk_related_locations_location_id 
    FOREIGN KEY (location_id) 
    REFERENCES locations(id) 
    ON DELETE CASCADE;

ALTER TABLE related_locations 
    ADD CONSTRAINT fk_related_locations_geo_location_id 
    FOREIGN KEY (geo_location_id) 
    REFERENCES geo_locations(id) 
    ON DELETE CASCADE;

-- Location Directions
CREATE TABLE IF NOT EXISTS location_directions (
    location_id BIGSERIAL NOT NULL,
    display_text_id BIGSERIAL NOT NULL
);

ALTER TABLE location_directions 
    ADD CONSTRAINT fk_location_directions_location_id 
    FOREIGN KEY (location_id) 
    REFERENCES locations(id) 
    ON DELETE CASCADE;

ALTER TABLE location_directions 
    ADD CONSTRAINT fk_location_directions_display_text_id
    FOREIGN KEY (display_text_id) 
    REFERENCES display_texts(id) 
    ON DELETE CASCADE;

-- Location Images
CREATE TABLE IF NOT EXISTS location_images (
    location_id BIGSERIAL NOT NULL,
    image_id BIGSERIAL NOT NULL
);

ALTER TABLE location_images
    ADD CONSTRAINT fk_location_images_location_id 
    FOREIGN KEY (location_id) 
    REFERENCES locations(id) 
    ON DELETE CASCADE;

ALTER TABLE location_images 
    ADD CONSTRAINT fk_location_images_image_id
    FOREIGN KEY (image_id) 
    REFERENCES images(id) 
    ON DELETE CASCADE;

-- Location Facilities
CREATE TABLE IF NOT EXISTS location_facilities (
    location_id BIGSERIAL NOT NULL,
    facility_id BIGSERIAL NOT NULL
);

ALTER TABLE location_facilities 
    ADD CONSTRAINT fk_location_facilities_location_id 
    FOREIGN KEY (location_id) 
    REFERENCES locations(id) 
    ON DELETE CASCADE;

ALTER TABLE location_facilities 
    ADD CONSTRAINT fk_location_facilities_facility_id
    FOREIGN KEY (facility_id) 
    REFERENCES facilities(id) 
    ON DELETE CASCADE;
