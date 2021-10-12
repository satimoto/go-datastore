package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

type EvsePayload struct {
	Evse  db.Evse
	Error error
}

type EvsesPayload struct {
	Evses []db.Evse
	Error error
}

func (r *MockRepositoryService) GetEvseByUid(ctx context.Context, uid string) (db.Evse, error) {
	if len(r.getEvseByUidPayload) == 0 {
		return db.Evse{}, nil
	}

	response := r.getEvseByUidPayload[0]
	r.getEvseByUidPayload = r.getEvseByUidPayload[1:]
	return response.Evse, response.Error

}

func (r *MockRepositoryService) ListEvses(ctx context.Context, locationID int64) ([]db.Evse, error) {
	if len(r.listEvsesPayload) == 0 {
		return []db.Evse{}, nil
	}

	response := r.listEvsesPayload[0]
	r.listEvsesPayload = r.listEvsesPayload[1:]
	return response.Evses, response.Error
}

func (r *MockRepositoryService) SetGetEvseByUidPayload(response EvsePayload) {
	r.getEvseByUidPayload = append(r.getEvseByUidPayload, response)
}

func (r *MockRepositoryService) SetListEvsesPayload(response EvsesPayload) {
	r.listEvsesPayload = append(r.listEvsesPayload, response)
}
