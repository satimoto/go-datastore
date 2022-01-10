-- name: CreateEmailSubscription :one
INSERT INTO email_subscriptions (
    email,
    locale,
    verification_code,
    unsubscribe_code,
    is_verified, 
    created_date
  ) VALUES ($1, $2, $3, $4, $5, $6)
  RETURNING *;

-- name: DeleteEmailSubscription :exec
DELETE FROM email_subscriptions
  WHERE id = $1;

-- name: GetEmailSubscriptionByEmail :one
SELECT * FROM email_subscriptions
  WHERE email = $1;

-- name: UpdateEmailSubscription :one
UPDATE email_subscriptions SET (
    email, 
    locale, 
    verification_code, 
    is_verified
  ) = ($2, $3, $4, $5)
  WHERE id = $1
  RETURNING *;
