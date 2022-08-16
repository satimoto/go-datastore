-- name: CreatePsbtFundingState :one
INSERT INTO psbt_funding_states (
    node_id,
    base_psbt,
    psbt, 
    expiry_date
  ) VALUES ($1, $2, $3, $4)
  RETURNING *;

-- name: GetPsbtFundingState :one
SELECT * FROM psbt_funding_states
  WHERE node_id = $1 AND funded_psbt is null;

-- name: ListPsbtFundingStates :many
SELECT * FROM psbt_funding_states
  WHERE node_id = $1 AND funded_psbt is null
  ORDER BY id;

-- name: UpdatePsbtFundingState :one
UPDATE psbt_funding_states SET (
    psbt,
    funded_psbt
  ) = ($2, $3)
  WHERE id = $1
  RETURNING *;
