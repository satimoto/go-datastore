-- name: CreateCommandReservation :one
INSERT INTO command_reservations (
    status,
    token_id,
    expiry_date,
    location_id,
    evse_uid
  ) VALUES ($1, $2, $3, $4, $5)
  RETURNING *;

-- name: GetCommandReservation :one
SELECT * FROM command_reservations
  WHERE id = $1;

-- name: UpdateCommandReservation :one
UPDATE command_reservations SET (
    status,
    expiry_date,
    evse_uid
  ) = ($2, $3, $4)
  WHERE id = $1
  RETURNING *;
