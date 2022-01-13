-- name: SetEvseCapability :exec
INSERT INTO evse_capabilities (evse_id, capability_id)
  VALUES ($1, $2);

-- name: UnsetEvseCapabilities :exec
DELETE FROM evse_capabilities ec
  WHERE ec.evse_id == $1;

-- name: ListEvseCapabilities :many
SELECT ca.* FROM capabilities ca
  INNER JOIN evse_capabilities ec ON ec.capability_id == ca.id
  WHERE ec.evse_id == $1
  ORDER BY ca.id;
