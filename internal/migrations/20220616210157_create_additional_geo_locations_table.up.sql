-- Geo Locations
ALTER TABLE IF EXISTS geo_locations 
    DROP IF EXISTS name;

-- Additional Geo Locations
CREATE TABLE IF NOT EXISTS additional_geo_locations (
    location_id     BIGINT NOT NULL,
    display_text_id BIGINT,
    latitude        TEXT NOT NULL,
    latitude_float  FLOAT NOT NULL,
    longitude       TEXT NOT NULL,
    longitude_float FLOAT NOT NULL
);

ALTER TABLE additional_geo_locations 
    ADD CONSTRAINT fk_additional_geo_locations_location_id 
    FOREIGN KEY (location_id) 
    REFERENCES locations(id) 
    ON DELETE CASCADE;

ALTER TABLE additional_geo_locations 
    ADD CONSTRAINT fk_additional_geo_locations_display_text_id
    FOREIGN KEY (display_text_id) 
    REFERENCES display_texts(id) 
    ON DELETE SET NULL;

-- Related Locations
ALTER TABLE IF EXISTS related_locations 
    DROP CONSTRAINT IF EXISTS fk_related_locations_geo_location_id;

ALTER TABLE IF EXISTS related_locations 
    DROP CONSTRAINT IF EXISTS fk_related_locations_location_id;

DROP TABLE IF EXISTS related_locations;
