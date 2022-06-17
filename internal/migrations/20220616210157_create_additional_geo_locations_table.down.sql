-- Related Locations
CREATE TABLE IF NOT EXISTS related_locations (
    location_id     BIGINT NOT NULL,
    geo_location_id BIGINT NOT NULL
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

-- Additional Geo Locations
ALTER TABLE IF EXISTS additional_geo_locations 
    DROP CONSTRAINT IF EXISTS fk_additional_geo_locations_display_text_id;

ALTER TABLE IF EXISTS additional_geo_locations 
    DROP CONSTRAINT IF EXISTS fk_additional_geo_locations_location_id;

DROP TABLE IF EXISTS additional_geo_locations;

-- Geo Locations
ALTER TABLE geo_locations 
    ADD COLUMN name;
