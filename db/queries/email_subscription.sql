-- name: CreateEmailSubscription :one
INSERT INTO email_subscriptions (
    email,
    is_verified, 
    created_date
  ) VALUES ($1, $2, $3)
  RETURNING *;

-- name: GetEmailSubscriptionByEmail :one
SELECT * FROM email_subscriptions
  WHERE email = $1;
