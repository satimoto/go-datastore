package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

type ConnectorMockData struct {
	Connector db.Connector
	Error     error
}

type ConnectorsMockData struct {
	Connectors []db.Connector
	Error      error
}

func (r *MockRepositoryService) CreateConnector(ctx context.Context, arg db.CreateConnectorParams) (db.Connector, error) {
	r.createConnectorMockData = append(r.createConnectorMockData, arg)
	return db.Connector{}, nil
}

func (r *MockRepositoryService) GetConnectorByUid(ctx context.Context, arg db.GetConnectorByUidParams) (db.Connector, error) {
	if len(r.getConnectorByUidMockData) == 0 {
		return db.Connector{}, nil
	}

	response := r.getConnectorByUidMockData[0]
	r.getConnectorByUidMockData = r.getConnectorByUidMockData[1:]
	return response.Connector, response.Error
}

func (r *MockRepositoryService) ListConnectors(ctx context.Context, id int64) ([]db.Connector, error) {
	if len(r.listConnectorsMockData) == 0 {
		return []db.Connector{}, nil
	}

	response := r.listConnectorsMockData[0]
	r.listConnectorsMockData = r.listConnectorsMockData[1:]
	return response.Connectors, response.Error
}

func (r *MockRepositoryService) UpdateConnectorByUid(ctx context.Context, arg db.UpdateConnectorByUidParams) (db.Connector, error) {
	if len(r.updateConnectorByUidMockData) == 0 {
		return db.Connector{}, nil
	}

	response := r.updateConnectorByUidMockData[0]
	r.updateConnectorByUidMockData = r.updateConnectorByUidMockData[1:]
	return response.Connector, response.Error
}

func (r *MockRepositoryService) GetCreateConnectorMockData() (db.CreateConnectorParams, error) {
	if len(r.createConnectorMockData) == 0 {
		return db.CreateConnectorParams{}, ErrorNotFound()
	}

	response := r.createConnectorMockData[0]
	r.createConnectorMockData = r.createConnectorMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetGetConnectorByUidMockData(response ConnectorMockData) {
	r.getConnectorByUidMockData = append(r.getConnectorByUidMockData, response)
}

func (r *MockRepositoryService) SetListConnectorsMockData(response ConnectorsMockData) {
	r.listConnectorsMockData = append(r.listConnectorsMockData, response)
}

func (r *MockRepositoryService) SetUpdateConnectorByUidMockData(response ConnectorMockData) {
	r.updateConnectorByUidMockData = append(r.updateConnectorByUidMockData, response)
}
