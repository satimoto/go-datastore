-- name: CreateCurrency :one
INSERT INTO currencies (
    name,
    code,
    numeric_code,
    decimals,
    symbol
  ) VALUES ($1, $2, $3, $4, $5)
  RETURNING *;

-- name: GetCurrencyByCode :one
SELECT * FROM currencies
  WHERE code = $1;
