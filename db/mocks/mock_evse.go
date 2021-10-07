package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

type EvsesResponse struct {
	Evses []db.Evse
	Error error
}

func (r *MockRepositoryService) ListEvses(ctx context.Context, locationID int64) ([]db.Evse, error) {
	return r.listEvsesResponse.Evses, r.listEvsesResponse.Error
}

func (r *MockRepositoryService) SetListEvsesResponse(response EvsesResponse) {
	r.listEvsesResponse = response
}
