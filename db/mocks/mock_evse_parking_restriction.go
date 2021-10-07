package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

func (r *MockRepositoryService) ListEvseParkingRestrictions(ctx context.Context, evseID int64) ([]db.ParkingRestriction, error) {
	return r.listEvseParkingRestrictionsResponse.ParkingRestrictions, r.listEvseParkingRestrictionsResponse.Error
}

func (r *MockRepositoryService) SetListEvseParkingRestrictionsResponse(response ParkingRestrictionsResponse) {
	r.listEvseParkingRestrictionsResponse = response
}
