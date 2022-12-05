package party

import (
	"context"

	"github.com/satimoto/go-datastore/pkg/db"
)

type PartyRepository interface {
	CreateParty(ctx context.Context, arg db.CreatePartyParams) (db.Party, error)
	GetParty(ctx context.Context, id int64) (db.Party, error)
	GetPartyByCredential(ctx context.Context, arg db.GetPartyByCredentialParams) (db.Party, error)
	UpdateParty(ctx context.Context, arg db.UpdatePartyParams) (db.Party, error)
	UpdatePartyByCredential(ctx context.Context, arg db.UpdatePartyByCredentialParams) (db.Party, error)
}

func NewRepository(repositoryService *db.RepositoryService) PartyRepository {
	return PartyRepository(repositoryService)
}
