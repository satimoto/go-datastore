-- name: SetCdrChargingPeriod :exec
INSERT INTO cdr_charging_periods (
    cdr_id, 
    charging_period_id
  ) VALUES ($1, $2);

-- name: DeleteCdrChargingPeriods :exec
DELETE FROM charging_periods cp
  USING cdr_charging_periods scp
  WHERE scp.charging_period_id == cp.id AND scp.cdr_id == $1;

-- name: ListCdrChargingPeriods :many
SELECT cp.* FROM charging_periods cp
  INNER JOIN cdr_charging_periods scp ON scp.charging_period_id == cp.id
  WHERE scp.cdr_id == $1
  ORDER BY cp.id;
