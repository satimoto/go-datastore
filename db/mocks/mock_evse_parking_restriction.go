package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

func (r *MockRepositoryService) ListEvseParkingRestrictions(ctx context.Context, evseID int64) ([]db.ParkingRestriction, error) {
	if len(r.listEvseParkingRestrictionsResponse) == 0 {
		return []db.ParkingRestriction{}, nil
	}

	response := r.listEvseParkingRestrictionsResponse[0]
	r.listEvseParkingRestrictionsResponse = r.listEvseParkingRestrictionsResponse[1:]
	return response.ParkingRestrictions, response.Error
}

func (r *MockRepositoryService) SetListEvseParkingRestrictionsResponse(response ParkingRestrictionsResponse) {
	r.listEvseParkingRestrictionsResponse = append(r.listEvseParkingRestrictionsResponse, response)
}
