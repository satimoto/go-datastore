-- name: SetTokenAuthorizationConnector :exec
INSERT INTO token_authorization_connectors (token_authorization_id, connector_uid)
  VALUES ($1, $2);

-- name: ListTokenAuthorizationConnectors :many
SELECT e.* FROM connectors e
  INNER JOIN token_authorization_connectors tac ON tac.connector_uid == e.uid
  WHERE tac.token_authorization_id == $1
  ORDER BY e.id;
