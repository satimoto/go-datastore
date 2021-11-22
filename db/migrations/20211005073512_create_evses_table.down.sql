-- Status Schedules
ALTER TABLE IF EXISTS status_schedules 
    DROP CONSTRAINT IF EXISTS fk_status_schedules_evse_id;

DROP TABLE IF EXISTS status_schedules;

-- Evse Parking Restrictions
ALTER TABLE IF EXISTS evse_parking_restrictions 
    DROP CONSTRAINT IF EXISTS fk_evse_parking_restrictions_parking_restriction_id;

ALTER TABLE IF EXISTS evse_parking_restrictions 
    DROP CONSTRAINT IF EXISTS fk_evse_parking_restrictions_evse_id;

DROP TABLE IF EXISTS evse_parking_restrictions;

-- Evse Capabilities
ALTER TABLE IF EXISTS evse_capabilities 
    DROP CONSTRAINT IF EXISTS fk_evse_capabilities_capability_id;

ALTER TABLE IF EXISTS evse_capabilities 
    DROP CONSTRAINT IF EXISTS fk_evse_capabilities_evse_id;

DROP TABLE IF EXISTS evse_capabilities;

-- Evse Images
ALTER TABLE IF EXISTS evse_images 
    DROP CONSTRAINT IF EXISTS fk_evse_images_image_id;

ALTER TABLE IF EXISTS evse_images 
    DROP CONSTRAINT IF EXISTS fk_evse_images_location_id;

DROP TABLE IF EXISTS evse_images;

-- Evse Directions
ALTER TABLE IF EXISTS evse_directions 
    DROP CONSTRAINT IF EXISTS fk_evse_directions_display_text_id;

ALTER TABLE IF EXISTS evse_directions 
    DROP CONSTRAINT IF EXISTS fk_evse_directions_evse_id;

DROP TABLE IF EXISTS evse_directions;

-- Connectors
ALTER TABLE IF EXISTS connectors 
    DROP CONSTRAINT IF EXISTS fk_connectors_evse_id;

ALTER TABLE IF EXISTS connectors 
    DROP CONSTRAINT IF EXISTS uq_connectors_evse_uid;

DROP INDEX IF EXISTS idx_connectors_uid;

DROP TABLE IF EXISTS connectors;

DROP TYPE IF EXISTS power_type;

DROP TYPE IF EXISTS connector_format;

DROP TYPE IF EXISTS connector_type;

-- EVSEs
ALTER TABLE IF EXISTS evses 
    DROP CONSTRAINT IF EXISTS fk_evses_location_id;

DROP INDEX IF EXISTS idx_evses_uid;

DROP TABLE IF EXISTS evses;

DROP TABLE IF EXISTS parking_restrictions;

DROP TABLE IF EXISTS capabilities;

DROP TYPE IF EXISTS evse_status;
