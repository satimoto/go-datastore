-- name: CreateNode :one
INSERT INTO nodes (
    pubkey,
    addr,
    alias,
    color,
    commit_hash,
    version,
    channels,
    peers
  ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
  RETURNING *;

-- name: GetNode :one
SELECT * FROM nodes
  WHERE id = $1;

-- name: GetNodeByPubkey :one
SELECT * FROM nodes
  WHERE pubkey = $1;

-- name: ListNodes :many
SELECT * FROM nodes
  ORDER BY peers ASC;

-- name: UpdateNode :one
UPDATE nodes SET (
    addr,
    alias,
    color,
    commit_hash,
    version,
    channels,
    peers
  ) = ($2, $3, $4, $5, $6, $7, $8)
  WHERE id = $1
  RETURNING *;
