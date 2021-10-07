package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

type ConnectorsResponse struct {
	Connectors []db.Connector
	Error      error
}

func (r *MockRepositoryService) ListConnectors(ctx context.Context, id int64) ([]db.Connector, error) {
	return r.listConnectorsResponse.Connectors, r.listConnectorsResponse.Error
}

func (r *MockRepositoryService) SetListConnectorsResponse(response ConnectorsResponse) {
	r.listConnectorsResponse = response
}
