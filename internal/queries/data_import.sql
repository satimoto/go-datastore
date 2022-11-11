-- name: GetHtbTariffByName :one
SELECT * FROM htb_tariffs
  WHERE name = $1;

-- name: ListHtbTariffs :many
SELECT * FROM htb_tariffs;
