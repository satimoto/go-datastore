package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

func (r *MockRepositoryService) ListEvseDirections(ctx context.Context, evseID int64) ([]db.DisplayText, error) {
	if len(r.listEvseDirectionsResponse) == 0 {
		return []db.DisplayText{}, nil
	}

	response := r.listEvseDirectionsResponse[0]
	r.listEvseDirectionsResponse = r.listEvseDirectionsResponse[1:]
	return response.DisplayTexts, response.Error
}

func (r *MockRepositoryService) SetListEvseDirectionsResponse(response DisplayTextsResponse) {
	r.listEvseDirectionsResponse = append(r.listEvseDirectionsResponse, response)
}
