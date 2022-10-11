package cdr

import (
	"context"

	"github.com/satimoto/go-datastore/pkg/db"
)

type CdrRepository interface {
	CreateCdr(ctx context.Context, arg db.CreateCdrParams) (db.Cdr, error)
	DeleteCdrChargingPeriods(ctx context.Context, cdrID int64) error
	GetCdrByLastUpdated(ctx context.Context, arg db.GetCdrByLastUpdatedParams) (db.Cdr, error)
	GetCdrByUid(ctx context.Context, uid string) (db.Cdr, error)
	ListCdrChargingPeriods(ctx context.Context, cdrID int64) ([]db.ChargingPeriod, error)
	ListCdrsBySessionStatus(ctx context.Context, statuses []db.SessionStatusType) ([]db.Cdr, error)
	SetCdrChargingPeriod(ctx context.Context, arg db.SetCdrChargingPeriodParams) error
}

func NewRepository(repositoryService *db.RepositoryService) CdrRepository {
	return CdrRepository(repositoryService)
}
