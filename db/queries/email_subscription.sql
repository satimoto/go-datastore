-- name: CreateEmailSubscription :one
INSERT INTO email_subscriptions (
    email,
    code,
    is_verified, 
    created_date
  ) VALUES ($1, $2, $3, $4)
  RETURNING *;

-- name: GetEmailSubscriptionByEmail :one
SELECT * FROM email_subscriptions
  WHERE email = $1;

-- name: UpdateEmailSubscription :one
UPDATE email_subscriptions SET (
    email, 
    code, 
    is_verified
  ) = ($2, $3, $4)
  WHERE id = $1
  RETURNING *;
