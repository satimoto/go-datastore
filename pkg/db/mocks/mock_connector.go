package mocks

import (
	"context"
	"database/sql"

	"github.com/satimoto/go-datastore/pkg/db"
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

func (r *MockRepositoryService) DeleteConnectors(ctx context.Context, evseID int64) error {
	r.deleteConnectorsMockData = append(r.deleteConnectorsMockData, evseID)
	return nil
}

func (r *MockRepositoryService) GetConnector(ctx context.Context, id int64) (db.Connector, error) {
	if len(r.getConnectorMockData) == 0 {
		return db.Connector{}, ErrorNotFound()
	}

	response := r.getConnectorMockData[0]
	r.getConnectorMockData = r.getConnectorMockData[1:]
	return response.Connector, response.Error
}

func (r *MockRepositoryService) GetConnectorByConnectorId(ctx context.Context, connectorID sql.NullString) (db.Connector, error) {
	if len(r.getConnectorByConnectorIdMockData) == 0 {
		return db.Connector{}, ErrorNotFound()
	}

	response := r.getConnectorByConnectorIdMockData[0]
	r.getConnectorByConnectorIdMockData = r.getConnectorByConnectorIdMockData[1:]
	return response.Connector, response.Error
}

func (r *MockRepositoryService) GetConnectorByUid(ctx context.Context, arg db.GetConnectorByUidParams) (db.Connector, error) {
	if len(r.getConnectorByUidMockData) == 0 {
		return db.Connector{}, ErrorNotFound()
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
	r.updateConnectorByUidMockData = append(r.updateConnectorByUidMockData, arg)
	return db.Connector{}, nil
}

func (r *MockRepositoryService) GetCreateConnectorMockData() (db.CreateConnectorParams, error) {
	if len(r.createConnectorMockData) == 0 {
		return db.CreateConnectorParams{}, ErrorNotFound()
	}

	response := r.createConnectorMockData[0]
	r.createConnectorMockData = r.createConnectorMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetDeleteConnectorsMockData() (int64, error) {
	if len(r.deleteConnectorsMockData) == 0 {
		return 0, ErrorNotFound()
	}

	response := r.deleteConnectorsMockData[0]
	r.deleteConnectorsMockData = r.deleteConnectorsMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetUpdateConnectorByUidMockData() (db.UpdateConnectorByUidParams, error) {
	if len(r.updateConnectorByUidMockData) == 0 {
		return db.UpdateConnectorByUidParams{}, ErrorNotFound()
	}

	response := r.updateConnectorByUidMockData[0]
	r.updateConnectorByUidMockData = r.updateConnectorByUidMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetGetConnectorMockData(response ConnectorMockData) {
	r.getConnectorMockData = append(r.getConnectorMockData, response)
}

func (r *MockRepositoryService) SetGetConnectorByConnectorIdMockData(response ConnectorMockData) {
	r.getConnectorByConnectorIdMockData = append(r.getConnectorByConnectorIdMockData, response)
}
func (r *MockRepositoryService) SetGetConnectorByUidMockData(response ConnectorMockData) {
	r.getConnectorByUidMockData = append(r.getConnectorByUidMockData, response)
}

func (r *MockRepositoryService) SetListConnectorsMockData(response ConnectorsMockData) {
	r.listConnectorsMockData = append(r.listConnectorsMockData, response)
}
