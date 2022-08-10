-- name: CreateCountryAccount :one
INSERT INTO country_accounts (
    country,
    tax_percent
  ) VALUES ($1, $2)
  RETURNING *;

-- name: GetCountryAccountByCountry :one
SELECT * FROM country_accounts
  WHERE country = $1;
