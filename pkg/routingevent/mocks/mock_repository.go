package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	"github.com/satimoto/go-datastore/pkg/routingevent"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) routingevent.RoutingEventRepository {
	return routingevent.RoutingEventRepository(repositoryService)
}
