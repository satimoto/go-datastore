package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	"github.com/satimoto/go-datastore/pkg/session"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) session.SessionRepository {
	return session.SessionRepository(repositoryService)
}
