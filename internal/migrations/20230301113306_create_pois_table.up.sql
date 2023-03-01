-- Pois
CREATE TABLE IF NOT EXISTS pois (
    id               BIGSERIAL PRIMARY KEY,
    uid              TEXT NOT NULL,
    source           TEXT NOT NULL,
    name             TEXT NOT NULL,
    geom             GEOMETRY(POINT, 4326) NOT NULL,
    description      TEXT,
    address          TEXT,
    city             TEXT,
    postal_code      TEXT,
    tag_key          TEXT NOT NULL,
    tag_value        TEXT NOT NULL,
    -- tags          []poi_tags
    payment_on_chain BOOLEAN NOT NULL,
    payment_ln       BOOLEAN NOT NULL,
    payment_ln_tap   BOOLEAN NOT NULL,
    opening_times    TEXT,
    phone            TEXT,
    website          TEXT,
    last_updated     TIMESTAMP NOT NULL
);

CREATE INDEX idx_pois_uid ON pois (uid);

CREATE INDEX idx_pois_geom ON pois USING GIST(geom);

-- Tags
CREATE TABLE IF NOT EXISTS tags (
    id    BIGSERIAL PRIMARY KEY,
    key   TEXT NOT NULL,
    value TEXT NOT NULL
);

ALTER TABLE tags 
    ADD CONSTRAINT uq_tags_key_value UNIQUE (key, value);

-- Poi tags
CREATE TABLE IF NOT EXISTS poi_tags (
    poi_id BIGINT NOT NULL,
    tag_id BIGINT NOT NULL
);

ALTER TABLE poi_tags
    ADD CONSTRAINT fk_poi_tags_poi_id
    FOREIGN KEY (poi_id) 
    REFERENCES pois(id) 
    ON DELETE CASCADE;

ALTER TABLE poi_tags 
    ADD CONSTRAINT fk_poi_tags_tag_id
    FOREIGN KEY (tag_id) 
    REFERENCES tags(id) 
    ON DELETE CASCADE;
