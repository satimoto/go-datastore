package db

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

type Repository interface {
	//User
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteUser(ctx context.Context, id int64) error
	GetUser(ctx context.Context, id int64) (User, error)
	ListUsers(ctx context.Context) ([]User, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)

	// Node
	CreateNode(ctx context.Context, arg CreateNodeParams) (Node, error)
	GetNode(ctx context.Context, id int64) (Node, error)
}

type RepositoryService struct {
	*Queries
	*sql.DB
}

func NewRepositoryService(d *sql.DB) *RepositoryService {
	return &RepositoryService{
		Queries: New(d),
		DB:      d,
	}
}

func Open(dataSourceName string) (*sql.DB, error) {
	return sql.Open("postgres", dataSourceName)
}