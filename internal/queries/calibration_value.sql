-- name: CreateCalibrationValue :one
INSERT INTO calibration_values (
    calibration_id,
    nature,
    plain_data,
    signed_data
  ) VALUES ($1, $2, $3, $4)
  RETURNING *;

-- name: ListCalibrationValues :many
SELECT * FROM calibration_values
  WHERE calibration_id = $1
  ORDER BY id;
