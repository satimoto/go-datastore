-- name: CreateConnector :one
INSERT INTO connectors (
    evse_id, 
    uid, 
    connector_id, 
    standard, 
    format, 
    power_type, 
    voltage, 
    amperage, 
    wattage, 
    tariff_id, 
    terms_and_conditions, 
    last_updated)
  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
  RETURNING *;

-- name: DeleteConnector :exec
DELETE FROM connectors
  WHERE id = $1;

-- name: DeleteConnectorByUid :exec
DELETE FROM connectors
  WHERE evse_id = $1 AND uid = $2;

-- name: DeleteConnectors :exec
DELETE FROM connectors
  WHERE evse_id = $1;

-- name: GetConnector :one
SELECT * FROM connectors
  WHERE id = $1;

-- name: GetConnectorByConnectorId :one
SELECT * FROM connectors
  WHERE connector_id = $1;

-- name: GetConnectorByUid :one
SELECT * FROM connectors
  WHERE (@evse_id::bigint IS NULL or evse_id = @evse_id::bigint) AND uid = @uid::string
  LIMIT 1;

-- name: ListConnectors :many
SELECT * FROM connectors
  WHERE evse_id = $1
  ORDER BY id;

-- name: UpdateConnector :one
UPDATE connectors SET (
    connector_id,
    standard, 
    format, 
    power_type, 
    voltage, 
    amperage, 
    wattage, 
    tariff_id, 
    terms_and_conditions, 
    last_updated
  ) = ($2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
  WHERE id = $1
  RETURNING *;

-- name: UpdateConnectorByUid :one
UPDATE connectors SET (
    connector_id,
    standard, 
    format, 
    power_type, 
    voltage, 
    amperage, 
    wattage, 
    tariff_id, 
    terms_and_conditions, 
    last_updated
  ) = ($3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
  WHERE evse_id = $1 AND uid = $2
  RETURNING *;
