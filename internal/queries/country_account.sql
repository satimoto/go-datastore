-- name: CreateCountryAccount :one
INSERT INTO country_accounts (
    country,
    tax_percent,
    longitude,
    latitude,
    zoom
  ) VALUES ($1, $2, $3, $4, $5)
  RETURNING *;

-- name: GetCountryAccountByCountry :one
SELECT * FROM country_accounts
  WHERE country = $1;

-- name: ListCountryAccounts :many
SELECT * FROM country_accounts
  ORDER BY country;
