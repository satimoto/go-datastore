-- name: CreateConnector :one
INSERT INTO connectors (
    evse_id, 
    uid, 
    identifier, 
    standard, 
    format, 
    power_type, 
    voltage, 
    amperage, 
    wattage, 
    tariff_id, 
    terms_and_conditions, 
    publish,
    last_updated)
  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
  RETURNING *;

-- name: DeleteConnector :exec
DELETE FROM connectors
  WHERE id = $1;

-- name: DeleteConnectorByEvse :exec
DELETE FROM connectors
  WHERE evse_id = $1 AND uid = $2;

-- name: DeleteConnectors :exec
DELETE FROM connectors
  WHERE evse_id = $1;

-- name: GetConnector :one
SELECT * FROM connectors
  WHERE id = $1;

-- name: GetConnectorByIdentifier :one
SELECT * FROM connectors
  WHERE identifier = $1;

-- name: GetConnectorByEvse :one
SELECT * FROM connectors
  WHERE evse_id = $1 AND uid = $2;
 
-- name: ListConnectors :many
SELECT * FROM connectors
  WHERE evse_id = $1 AND publish = true
  ORDER BY id;

-- name: UpdateConnector :one
UPDATE connectors SET (
    identifier,
    standard, 
    format, 
    power_type, 
    voltage, 
    amperage, 
    wattage, 
    tariff_id, 
    terms_and_conditions, 
    publish,
    last_updated
  ) = ($2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
  WHERE id = $1
  RETURNING *;

-- name: UpdateConnectorByEvse :one
UPDATE connectors SET (
    identifier,
    standard, 
    format, 
    power_type, 
    voltage, 
    amperage, 
    wattage, 
    tariff_id, 
    terms_and_conditions, 
    publish,
    last_updated
  ) = ($3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
  WHERE evse_id = $1 AND uid = $2
  RETURNING *;
