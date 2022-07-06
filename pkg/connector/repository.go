package connector

import (
	"context"
	"database/sql"

	"github.com/satimoto/go-datastore/pkg/db"
)

type ConnectorRepository interface {
	CreateConnector(ctx context.Context, arg db.CreateConnectorParams) (db.Connector, error)
	GetConnectorByConnectorId(ctx context.Context, connectorID sql.NullString) (db.Connector, error)
	GetConnectorByUid(ctx context.Context, arg db.GetConnectorByUidParams) (db.Connector, error)
	GetGeoLocation(ctx context.Context, id int64) (db.GeoLocation, error)
	GetEvse(ctx context.Context, id int64) (db.Evse, error)
	ListConnectors(ctx context.Context, evseID int64) ([]db.Connector, error)
	UpdateConnectorByUid(ctx context.Context, arg db.UpdateConnectorByUidParams) (db.Connector, error)
	UpdateEvseLastUpdated(ctx context.Context, arg db.UpdateEvseLastUpdatedParams) error
	UpdateLocationLastUpdated(ctx context.Context, arg db.UpdateLocationLastUpdatedParams) error
}

func NewRepository(repositoryService *db.RepositoryService) ConnectorRepository {
	return ConnectorRepository(repositoryService)
}
