-- name: SetTokenAuthorizationConnector :exec
INSERT INTO token_authorization_connectors (token_authorization_id, connector_id)
  VALUES ($1, $2);

-- name: ListTokenAuthorizationConnectors :many
SELECT c.* FROM connectors c
  INNER JOIN token_authorization_connectors tac ON tac.connector_id = c.id
  WHERE tac.token_authorization_id = $1
  ORDER BY c.id;
