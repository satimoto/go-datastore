package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

func (r *MockRepositoryService) ListEvseCapabilities(ctx context.Context, evseID int64) ([]db.Capability, error) {
	if len(r.listEvseCapabilitiesPayload) == 0 {
		return []db.Capability{}, nil
	}

	response := r.listEvseCapabilitiesPayload[0]
	r.listEvseCapabilitiesPayload = r.listEvseCapabilitiesPayload[1:]
	return response.Capabilities, response.Error
}

func (r *MockRepositoryService) SetListEvseCapabilitiesPayload(response CapabilitiesPayload) {
	r.listEvseCapabilitiesPayload = append(r.listEvseCapabilitiesPayload, response)
}
