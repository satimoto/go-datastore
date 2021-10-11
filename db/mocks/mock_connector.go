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
	if len(r.listConnectorsResponse) == 0 {
		return []db.Connector{}, nil
	}

	response := r.listConnectorsResponse[0]
	r.listConnectorsResponse = r.listConnectorsResponse[1:]
	return response.Connectors, response.Error
}

func (r *MockRepositoryService) SetListConnectorsResponse(response ConnectorsResponse) {
	r.listConnectorsResponse = append(r.listConnectorsResponse, response)
}
