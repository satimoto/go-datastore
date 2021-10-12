package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

func (r *MockRepositoryService) ListEvseDirections(ctx context.Context, evseID int64) ([]db.DisplayText, error) {
	if len(r.listEvseDirectionsPayload) == 0 {
		return []db.DisplayText{}, nil
	}

	response := r.listEvseDirectionsPayload[0]
	r.listEvseDirectionsPayload = r.listEvseDirectionsPayload[1:]
	return response.DisplayTexts, response.Error
}

func (r *MockRepositoryService) SetListEvseDirectionsPayload(response DisplayTextsPayload) {
	r.listEvseDirectionsPayload = append(r.listEvseDirectionsPayload, response)
}
