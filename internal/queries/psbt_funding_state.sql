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
  WHERE id = $1;

-- name: GetUnfundedPsbtFundingState :one
SELECT * FROM psbt_funding_states
  WHERE node_id = $1 AND funded_psbt is null AND not is_failed;

-- name: ListUnfundedPsbtFundingStates :many
SELECT * FROM psbt_funding_states
  WHERE node_id = $1 AND funded_psbt is null AND not is_failed
  ORDER BY id;

-- name: UpdatePsbtFundingState :one
UPDATE psbt_funding_states SET (
    psbt,
    funded_psbt,
    signed_psbt,
    signed_tx,
    is_failed
  ) = ($2, $3, $4, $5, $6)
  WHERE id = $1
  RETURNING *;
