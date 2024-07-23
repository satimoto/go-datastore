
-- Poi tags
ALTER TABLE IF EXISTS poi_tags 
    DROP CONSTRAINT IF EXISTS fk_poi_tags_tag_id;

ALTER TABLE IF EXISTS poi_tags 
    DROP CONSTRAINT IF EXISTS fk_poi_tags_poi_id;

DROP TABLE IF EXISTS poi_tags;

-- Tags
ALTER TABLE IF EXISTS tags 
    DROP CONSTRAINT IF EXISTS uq_tags_key_value;

DROP TABLE IF EXISTS tags;

-- Pois
DROP INDEX IF EXISTS idx_pois_geom;

DROP INDEX IF EXISTS idx_pois_uid;

DROP TABLE IF EXISTS pois;
