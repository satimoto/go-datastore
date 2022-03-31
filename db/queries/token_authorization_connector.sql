-- name: SetTokenAuthorizationConnector :exec
INSERT INTO token_authorization_connectors (token_authorization_id, connector_id)
  VALUES ($1, $2);

-- name: ListTokenAuthorizationConnectors :many
SELECT e.* FROM connectors e
  INNER JOIN token_authorization_connectors tac ON tac.connector_id == e.id
  WHERE tac.token_authorization_id == $1
  ORDER BY e.id;
