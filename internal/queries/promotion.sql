-- name: CreatePromotion :one
INSERT INTO promotions (
    code,
    description,
    start_date,
    end_date
  ) VALUES ($1, $2, $3, $4)
  RETURNING *;

-- name: GetPromotion :one
SELECT * FROM promotions
  WHERE id = $1;

-- name: GetPromotionByCode :one
SELECT * FROM promotions
  WHERE code = $1;
