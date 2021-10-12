package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

func (r *MockRepositoryService) ListLocationFacilities(ctx context.Context, locationID int64) ([]db.Facility, error) {
	if len(r.listLocationFacilitiesPayload) == 0 {
		return []db.Facility{}, nil
	}

	response := r.listLocationFacilitiesPayload[0]
	r.listLocationFacilitiesPayload = r.listLocationFacilitiesPayload[1:]
	return response.Facilities, response.Error
}

func (r *MockRepositoryService) SetListLocationFacilitiesPayload(response FacilitiesPayload) {
	r.listLocationFacilitiesPayload = append(r.listLocationFacilitiesPayload, response)
}
