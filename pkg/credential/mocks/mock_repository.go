package mocks

import (
	"github.com/satimoto/go-datastore/pkg/credential"
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) credential.CredentialRepository {
	return credential.CredentialRepository(repositoryService)
}
