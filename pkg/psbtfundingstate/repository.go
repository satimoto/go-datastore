package psbtfundingstate

import (
	"context"

	"github.com/satimoto/go-datastore/pkg/db"
)

type PsbtFundingStateRepository interface {
	CreatePsbtFundingState(ctx context.Context, arg db.CreatePsbtFundingStateParams) (db.PsbtFundingState, error)
	GetPsbtFundingState(ctx context.Context, nodeID int64) (db.PsbtFundingState, error)
	ListPsbtFundingStates(ctx context.Context, nodeID int64) ([]db.PsbtFundingState, error)
	ListPsbtFundingStateChannelRequests(ctx context.Context, psbtFundingStateID int64) ([]db.ChannelRequest, error)
	SetPsbtFundingStateChannelRequest(ctx context.Context, arg db.SetPsbtFundingStateChannelRequestParams) error
	UpdatePsbtFundingState(ctx context.Context, arg db.UpdatePsbtFundingStateParams) (db.PsbtFundingState, error)
}

func NewRepository(repositoryService *db.RepositoryService) PsbtFundingStateRepository {
	return PsbtFundingStateRepository(repositoryService)
}
