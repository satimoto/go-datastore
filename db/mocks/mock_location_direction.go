package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

func (r *MockRepositoryService) ListLocationDirections(ctx context.Context, locationID int64) ([]db.DisplayText, error) {
	if len(r.listLocationDirectionsPayload) == 0 {
		return []db.DisplayText{}, nil
	}

	response := r.listLocationDirectionsPayload[0]
	r.listLocationDirectionsPayload = r.listLocationDirectionsPayload[1:]
	return response.DisplayTexts, response.Error
}

func (r *MockRepositoryService) SetListLocationDirectionsPayload(response DisplayTextsPayload) {
	r.listLocationDirectionsPayload = append(r.listLocationDirectionsPayload, response)
}
