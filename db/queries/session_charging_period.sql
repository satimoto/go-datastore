-- name: SetSessionChargingPeriod :exec
INSERT INTO session_charging_periods (
    session_id, 
    charging_period_id
  ) VALUES ($1, $2);

-- name: DeleteSessionChargingPeriods :exec
DELETE FROM charging_periods cp
  USING session_charging_periods scp
  WHERE scp.charging_period_id == cp.id AND scp.session_id == $1;

-- name: ListSessionChargingPeriods :many
SELECT cp.* FROM charging_periods cp
  INNER JOIN session_charging_periods scp ON scp.charging_period_id == cp.id
  WHERE scp.session_id == $1
  ORDER BY cp.id;
