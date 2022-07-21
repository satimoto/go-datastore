-- name: CreateNode :one
INSERT INTO nodes (
    pubkey,
    node_addr,
    lsp_addr,
    alias,
    color,
    commit_hash,
    version,
    channels,
    peers,
    is_active
  ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
  RETURNING *;

-- name: GetNode :one
SELECT * FROM nodes
  WHERE id = $1;

-- name: GetNodeByPubkey :one
SELECT * FROM nodes
  WHERE pubkey = $1;

-- name: GetNodeByUserID :one
SELECT n.* FROM nodes n
  INNER JOIN users u ON u.node_id = n.id
  WHERE u.id = $1;

-- name: ListNodes :many
SELECT * FROM nodes
  ORDER BY peers ASC;

-- name: ListActiveNodes :many
SELECT * FROM nodes
  WHERE is_active = true
  ORDER BY peers ASC;

-- name: UpdateNode :one
UPDATE nodes SET (
    node_addr,
    lsp_addr,
    alias,
    color,
    commit_hash,
    version,
    channels,
    peers,
    is_active
  ) = ($2, $3, $4, $5, $6, $7, $8, $9, $10)
  WHERE id = $1
  RETURNING *;
