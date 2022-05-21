package displaytext

import (
	"context"

	"github.com/satimoto/go-datastore/pkg/db"
)

type DisplayTextRepository interface {
	CreateDisplayText(ctx context.Context, arg db.CreateDisplayTextParams) (db.DisplayText, error)
}

func NewRepository(repositoryService *db.RepositoryService) DisplayTextRepository {
	return DisplayTextRepository(repositoryService)
}
