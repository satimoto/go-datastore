-- name: SetEvseParkingRestriction :exec
INSERT INTO evse_parking_restrictions (evse_id, parking_restriction_id)
  VALUES ($1, $2);

-- name: UnsetEvseParkingRestrictions :exec
DELETE FROM evse_parking_restrictions ep
  WHERE ep.evse_id = $1;

-- name: ListEvseParkingRestrictions :many
SELECT pr.* FROM parking_restrictions pr
  INNER JOIN evse_parking_restrictions ep ON ep.parking_restriction_id = pr.id
  WHERE ep.evse_id = $1
  ORDER BY pr.id;
