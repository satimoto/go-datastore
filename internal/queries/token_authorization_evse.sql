-- name: SetTokenAuthorizationEvse :exec
INSERT INTO token_authorization_evses (token_authorization_id, evse_id)
  VALUES ($1, $2);

-- name: ListTokenAuthorizationEvses :many
SELECT e.* FROM evses e
  INNER JOIN token_authorization_evses tae ON tae.evse_id = e.id
  WHERE tae.token_authorization_id = $1
  ORDER BY e.id;
