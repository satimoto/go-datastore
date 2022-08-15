-- name: CreatePsbtFundingState :one
INSERT INTO psbt_funding_states (
    base_psbt,
    psbt, 
    expiry_date
  ) VALUES ($1, $2, $3)
  RETURNING *;

-- name: ListPsbtFundingStates :many
SELECT * FROM psbt_funding_states
  WHERE funded_psbt is null
  ORDER BY id;

-- name: UpdatePsbtFundingState :one
UPDATE psbt_funding_states SET (
    psbt,
    funded_psbt
  ) = ($2, $3)
  WHERE id = $1
  RETURNING *;
