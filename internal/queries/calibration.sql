-- name: CreateCalibration :one
INSERT INTO calibrations (
    encoding_method,
    encoding_method_version,
    public_key,
    url
  ) VALUES ($1, $2, $3, $4)
  RETURNING *;

-- name: GetCalibration :one
SELECT * FROM calibrations
  WHERE id = $1;
