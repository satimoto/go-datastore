-- name: CreatePoi :one
INSERT INTO pois (
    uid,
    source,
    name,
    geom,
    description,
    address,
    city,
    postal_code,
    tag_key,
    tag_value,
    payment_on_chain,
    payment_ln,
    payment_ln_tap,
    opening_times,
    phone,
    website,
    last_updated
  ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)
  RETURNING *;

-- name: DeletePoiByUid :exec
DELETE FROM pois
  WHERE uid = $1;

-- name: GetPoi :one
SELECT * FROM pois
  WHERE id = $1;

-- name: GetPoiByLastUpdated :one
SELECT * FROM pois
  ORDER BY last_updated DESC
  LIMIT 1;

-- name: ListPoisByGeom :many
SELECT * FROM pois
  WHERE ST_Intersects(geom, ST_MakeEnvelope(@x_min::FLOAT, @y_min::FLOAT, @x_max::FLOAT, @y_max::FLOAT, 4326))
  LIMIT 100;

-- name: UpdatePoiByUid :one
UPDATE pois SET (
    name,
    geom,
    description,
    address,
    city,
    postal_code,
    tag_key,
    tag_value,
    payment_on_chain,
    payment_ln,
    payment_ln_tap,
    opening_times,
    phone,
    website,
    last_updated
  ) = ($2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)
  WHERE uid = $1
  RETURNING *;
