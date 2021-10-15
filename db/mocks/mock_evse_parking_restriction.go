package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

func (r *MockRepositoryService) ListEvseParkingRestrictions(ctx context.Context, evseID int64) ([]db.ParkingRestriction, error) {
	if len(r.listEvseParkingRestrictionsMockData) == 0 {
		return []db.ParkingRestriction{}, nil
	}

	response := r.listEvseParkingRestrictionsMockData[0]
	r.listEvseParkingRestrictionsMockData = r.listEvseParkingRestrictionsMockData[1:]
	return response.ParkingRestrictions, response.Error
}

func (r *MockRepositoryService) SetListEvseParkingRestrictionsMockData(response ParkingRestrictionsMockData) {
	r.listEvseParkingRestrictionsMockData = append(r.listEvseParkingRestrictionsMockData, response)
}
