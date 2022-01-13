-- Location Facilities
ALTER TABLE IF EXISTS location_facilities 
    DROP CONSTRAINT IF EXISTS fk_location_facilities_facility_id;

ALTER TABLE IF EXISTS location_facilities 
    DROP CONSTRAINT IF EXISTS fk_location_facilities_location_id;

DROP TABLE IF EXISTS location_facilities;

-- Location Images
ALTER TABLE IF EXISTS location_images 
    DROP CONSTRAINT IF EXISTS fk_location_images_image_id;

ALTER TABLE IF EXISTS location_images 
    DROP CONSTRAINT IF EXISTS fk_location_images_location_id;

DROP TABLE IF EXISTS location_images;

-- Location Directions
ALTER TABLE IF EXISTS location_directions 
    DROP CONSTRAINT IF EXISTS fk_location_directions_display_text_id;

ALTER TABLE IF EXISTS location_directions 
    DROP CONSTRAINT IF EXISTS fk_location_directions_location_id;

DROP TABLE IF EXISTS location_directions;

-- Related Locations
ALTER TABLE IF EXISTS related_locations 
    DROP CONSTRAINT IF EXISTS fk_related_locations_geo_location_id;

ALTER TABLE IF EXISTS related_locations 
    DROP CONSTRAINT IF EXISTS fk_related_locations_location_id;

DROP TABLE IF EXISTS related_locations;

-- Locations
ALTER TABLE IF EXISTS locations 
    DROP CONSTRAINT IF EXISTS fk_locations_opening_time_id;

ALTER TABLE IF EXISTS locations 
    DROP CONSTRAINT IF EXISTS fk_locations_energy_mix_id;

ALTER TABLE IF EXISTS locations 
    DROP CONSTRAINT IF EXISTS fk_locations_owner_id;

ALTER TABLE IF EXISTS locations 
    DROP CONSTRAINT IF EXISTS fk_locations_suboperator_id;

ALTER TABLE IF EXISTS locations 
    DROP CONSTRAINT IF EXISTS fk_locations_operator_id;

DROP INDEX IF EXISTS idx_locations_uid;

DROP TABLE IF EXISTS locations;

DROP TYPE IF EXISTS location_type;

-- Location Hours
ALTER TABLE IF EXISTS exceptional_periods 
    DROP CONSTRAINT IF EXISTS fk_exceptional_periods_opening_time_id;

DROP TABLE IF EXISTS exceptional_periods;

ALTER TABLE IF EXISTS regular_hours 
    DROP CONSTRAINT IF EXISTS fk_regular_hours_opening_time_id;

DROP TABLE IF EXISTS regular_hours;

DROP TABLE IF EXISTS opening_times;

DROP TYPE IF EXISTS period_type;

-- Energy Mix
ALTER TABLE IF EXISTS energy_sources 
    DROP CONSTRAINT IF EXISTS fk_energy_sources_energy_mix_id;

DROP TABLE IF EXISTS energy_sources;

ALTER TABLE IF EXISTS environmental_impacts 
    DROP CONSTRAINT IF EXISTS fk_environmental_impacts_energy_mix_id;

DROP TABLE IF EXISTS environmental_impacts;

DROP TABLE IF EXISTS energy_mixes;

DROP TYPE IF EXISTS environmental_impact_category;

DROP TYPE IF EXISTS energy_source_category;

-- Business Details
ALTER TABLE IF EXISTS business_details 
    DROP CONSTRAINT IF EXISTS fk_business_details_logo_id;

DROP TABLE IF EXISTS business_details;

-- Facilities
DROP TABLE IF EXISTS facilities;

-- Images
DROP TABLE IF EXISTS images;

DROP TYPE IF EXISTS image_category;

-- Geo Locations
DROP TABLE IF EXISTS geo_locations;

-- Display Texts
DROP TABLE IF EXISTS display_texts;
