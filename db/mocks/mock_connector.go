package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

type ConnectorPayload struct {
	Connector db.Connector
	Error     error
}

type ConnectorsPayload struct {
	Connectors []db.Connector
	Error      error
}

func (r *MockRepositoryService) GetConnectorByUid(ctx context.Context, arg db.GetConnectorByUidParams) (db.Connector, error) {
	if len(r.getConnectorByUidPayload) == 0 {
		return db.Connector{}, nil
	}

	response := r.getConnectorByUidPayload[0]
	r.getConnectorByUidPayload = r.getConnectorByUidPayload[1:]
	return response.Connector, response.Error
}

func (r *MockRepositoryService) ListConnectors(ctx context.Context, id int64) ([]db.Connector, error) {
	if len(r.listConnectorsPayload) == 0 {
		return []db.Connector{}, nil
	}

	response := r.listConnectorsPayload[0]
	r.listConnectorsPayload = r.listConnectorsPayload[1:]
	return response.Connectors, response.Error
}

func (r *MockRepositoryService) UpdateConnectorByUid(ctx context.Context, arg db.UpdateConnectorByUidParams) (db.Connector, error) {
	if len(r.updateConnectorByUidPayload) == 0 {
		return db.Connector{}, nil
	}

	response := r.updateConnectorByUidPayload[0]
	r.updateConnectorByUidPayload = r.updateConnectorByUidPayload[1:]
	return response.Connectors, response.Error
}

func (r *MockRepositoryService) SetGetConnectorByUidPayload(response ConnectorPayload) {
	r.getConnectorByUidPayload = append(r.getConnectorByUidPayload, response)
}

func (r *MockRepositoryService) SetListConnectorsPayload(response ConnectorsPayload) {
	r.listConnectorsPayload = append(r.listConnectorsPayload, response)
}

func (r *MockRepositoryService) SetUpdateConnectorByUidPayload(response ConnectorPayload) {
	r.updateConnectorByUidPayload = append(r.updateConnectorByUidPayload, response)
}
