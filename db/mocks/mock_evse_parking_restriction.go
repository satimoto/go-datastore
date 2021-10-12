package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

func (r *MockRepositoryService) ListEvseParkingRestrictions(ctx context.Context, evseID int64) ([]db.ParkingRestriction, error) {
	if len(r.listEvseParkingRestrictionsPayload) == 0 {
		return []db.ParkingRestriction{}, nil
	}

	response := r.listEvseParkingRestrictionsPayload[0]
	r.listEvseParkingRestrictionsPayload = r.listEvseParkingRestrictionsPayload[1:]
	return response.ParkingRestrictions, response.Error
}

func (r *MockRepositoryService) SetListEvseParkingRestrictionsPayload(response ParkingRestrictionsPayload) {
	r.listEvseParkingRestrictionsPayload = append(r.listEvseParkingRestrictionsPayload, response)
}
