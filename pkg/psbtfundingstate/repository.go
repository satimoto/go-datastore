package psbtfundingstate

import (
	"context"

	"github.com/satimoto/go-datastore/pkg/db"
)

type PsbtFundingStateRepository interface {
	CreatePsbtFundingState(ctx context.Context, arg db.CreatePsbtFundingStateParams) (db.PsbtFundingState, error)
	ListPsbtFundingStates(ctx context.Context) ([]db.PsbtFundingState, error)
	ListPsbtFundingStateChannelRequests(ctx context.Context, psbtFundingStateID int64) ([]db.ChannelRequest, error)
	SetPsbtFundingStateChannelRequest(ctx context.Context, arg db.SetPsbtFundingStateChannelRequestParams) error
	UpdatePsbtFundingState(ctx context.Context, arg db.UpdatePsbtFundingStateParams) (db.PsbtFundingState, error)
}

func NewRepository(repositoryService *db.RepositoryService) PsbtFundingStateRepository {
	return PsbtFundingStateRepository(repositoryService)
}
