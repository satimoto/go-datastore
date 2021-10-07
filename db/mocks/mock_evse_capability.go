package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

func (r *MockRepositoryService) ListEvseCapabilities(ctx context.Context, evseID int64) ([]db.Capability, error) {
	return r.listEvseCapabilitiesResponse.Capabilities, r.listEvseCapabilitiesResponse.Error
}

func (r *MockRepositoryService) SetListEvseCapabilitiesResponse(response CapabilitiesResponse) {
	r.listEvseCapabilitiesResponse = response
}
