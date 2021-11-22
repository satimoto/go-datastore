// Code generated by sqlc. DO NOT EDIT.
// source: email_subscription.sql

package db

import (
	"context"
	"time"
)

const createEmailSubscription = `-- name: CreateEmailSubscription :one
INSERT INTO email_subscriptions (
    email,
    is_verified, 
    created_date
  ) VALUES ($1, $2, $3)
  RETURNING id, email, is_verified, created_date
`

type CreateEmailSubscriptionParams struct {
	Email       string    `db:"email" json:"email"`
	IsVerified  bool      `db:"is_verified" json:"isVerified"`
	CreatedDate time.Time `db:"created_date" json:"createdDate"`
}

func (q *Queries) CreateEmailSubscription(ctx context.Context, arg CreateEmailSubscriptionParams) (EmailSubscription, error) {
	row := q.db.QueryRowContext(ctx, createEmailSubscription, arg.Email, arg.IsVerified, arg.CreatedDate)
	var i EmailSubscription
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.IsVerified,
		&i.CreatedDate,
	)
	return i, err
}

const getEmailSubscriptionByEmail = `-- name: GetEmailSubscriptionByEmail :one
SELECT id, email, is_verified, created_date FROM email_subscriptions
  WHERE email = $1
`

func (q *Queries) GetEmailSubscriptionByEmail(ctx context.Context, email string) (EmailSubscription, error) {
	row := q.db.QueryRowContext(ctx, getEmailSubscriptionByEmail, email)
	var i EmailSubscription
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.IsVerified,
		&i.CreatedDate,
	)
	return i, err
}
