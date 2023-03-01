-- name: CreateCommandReservation :one
INSERT INTO command_reservations (
    status,
    token_id,
    expiry_date,
    location_id,
    evse_uid,
    last_updated
  ) VALUES ($1, $2, $3, $4, $5, $6)
  RETURNING *;

-- name: GetCommandReservation :one
SELECT * FROM command_reservations
  WHERE id = $1;

-- name: UpdateCommandReservation :one
UPDATE command_reservations SET (
    status,
    expiry_date,
    evse_uid,
    last_updated
  ) = ($2, $3, $4, $5)
  WHERE id = $1
  RETURNING *;
