-- name: CreateReferral :one
INSERT INTO referrals (
    promotion_id,
    user_id,
    ip_address,
    last_updated
  ) VALUES ($1, $2, $3, $4)
  RETURNING *;

-- name: GetReferralByIpAddress :one
SELECT * FROM referrals
  WHERE ip_address = $1;
