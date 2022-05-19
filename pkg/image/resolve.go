package image

import (
	"context"

	"github.com/satimoto/go-datastore/pkg/db"
)

type ImageRepository interface {
	CreateImage(ctx context.Context, arg db.CreateImageParams) (db.Image, error)
	DeleteImage(ctx context.Context, id int64) error
	GetImage(ctx context.Context, id int64) (db.Image, error)
	UpdateImage(ctx context.Context, arg db.UpdateImageParams) (db.Image, error)
}

type ImageResolver struct {
	Repository ImageRepository
}

func NewResolver(repositoryService *db.RepositoryService) *ImageResolver {
	repo := ImageRepository(repositoryService)
	return &ImageResolver{repo}
}
