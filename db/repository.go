package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

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