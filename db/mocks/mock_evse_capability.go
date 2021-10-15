package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

func (r *MockRepositoryService) ListEvseCapabilities(ctx context.Context, evseID int64) ([]db.Capability, error) {
	if len(r.listEvseCapabilitiesMockData) == 0 {
		return []db.Capability{}, nil
	}

	response := r.listEvseCapabilitiesMockData[0]
	r.listEvseCapabilitiesMockData = r.listEvseCapabilitiesMockData[1:]
	return response.Capabilities, response.Error
}

func (r *MockRepositoryService) SetListEvseCapabilitiesMockData(response CapabilitiesMockData) {
	r.listEvseCapabilitiesMockData = append(r.listEvseCapabilitiesMockData, response)
}
