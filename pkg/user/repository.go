package user

import (
	"context"

	"github.com/satimoto/go-datastore/pkg/db"
)

type UserRepository interface {
	CreateUser(ctx context.Context, arg db.CreateUserParams) (db.User, error)
	GetUser(ctx context.Context, id int64) (db.User, error)
	GetUserByDeviceToken(ctx context.Context, deviceToken string) (db.User, error)
	GetUserByLinkingPubkey(ctx context.Context, linkingPubkey string) (db.User, error)
	GetUserBySessionID(ctx context.Context, id int64) (db.User, error)
	GetUserByTokenID(ctx context.Context, id int64) (db.User, error)
	UpdateUser(ctx context.Context, arg db.UpdateUserParams) (db.User, error)
}

func NewRepository(repositoryService *db.RepositoryService) UserRepository {
	return UserRepository(repositoryService)
}
