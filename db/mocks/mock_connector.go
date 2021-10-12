package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

type ConnectorResponse struct {
	Connector db.Connector
	Error     error
}

type ConnectorsResponse struct {
	Connectors []db.Connector
	Error      error
}

func (r *MockRepositoryService) GetConnectorByUid(ctx context.Context, arg db.GetConnectorByUidParams) (db.Connector, error) {
	if len(r.getConnectorByUidResponse) == 0 {
		return db.Connector{}, nil
	}

	response := r.getConnectorByUidResponse[0]
	r.getConnectorByUidResponse = r.getConnectorByUidResponse[1:]
	return response.Connector, response.Error
}

func (r *MockRepositoryService) ListConnectors(ctx context.Context, id int64) ([]db.Connector, error) {
	if len(r.listConnectorsResponse) == 0 {
		return []db.Connector{}, nil
	}

	response := r.listConnectorsResponse[0]
	r.listConnectorsResponse = r.listConnectorsResponse[1:]
	return response.Connectors, response.Error
}

func (r *MockRepositoryService) SetGetConnectorByUidResponse(response ConnectorResponse) {
	r.getConnectorByUidResponse = append(r.getConnectorByUidResponse, response)
}

func (r *MockRepositoryService) SetListConnectorsResponse(response ConnectorsResponse) {
	r.listConnectorsResponse = append(r.listConnectorsResponse, response)
}
