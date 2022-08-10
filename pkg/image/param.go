package image

import "github.com/satimoto/go-datastore/pkg/db"

func NewUpdateImageParams(image db.Image) db.UpdateImageParams {
	return db.UpdateImageParams(image)
}
