package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

func (r *MockRepositoryService) ListEvseCapabilities(ctx context.Context, evseID int64) ([]db.Capability, error) {
	if len(r.listEvseCapabilitiesResponse) == 0 {
		return []db.Capability{}, nil
	}

	response := r.listEvseCapabilitiesResponse[0]
	r.listEvseCapabilitiesResponse = r.listEvseCapabilitiesResponse[1:]
	return response.Capabilities, response.Error
}

func (r *MockRepositoryService) SetListEvseCapabilitiesResponse(response CapabilitiesResponse) {
	r.listEvseCapabilitiesResponse = append(r.listEvseCapabilitiesResponse, response)
}
