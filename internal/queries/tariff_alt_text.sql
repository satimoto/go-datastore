-- name: SetTariffAltText :exec
INSERT INTO tariff_alt_texts (tariff_id, display_text_id)
  VALUES ($1, $2);

-- name: DeleteTariffAltTexts :exec
DELETE FROM display_texts dt
  USING tariff_alt_texts tt
  WHERE tt.display_text_id = dt.id AND tt.tariff_id = $1;

-- name: ListTariffAltTexts :many
SELECT dt.* FROM display_texts dt
  INNER JOIN tariff_alt_texts tt ON tt.display_text_id = dt.id
  WHERE tt.tariff_id = $1
  ORDER BY dt.id;
