package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

func (r *MockRepositoryService) ListLocationDirections(ctx context.Context, locationID int64) ([]db.DisplayText, error) {
	if len(r.listLocationDirectionsResponse) == 0 {
		return []db.DisplayText{}, nil
	}

	response := r.listLocationDirectionsResponse[0]
	r.listLocationDirectionsResponse = r.listLocationDirectionsResponse[1:]
	return response.DisplayTexts, response.Error
}

func (r *MockRepositoryService) SetListLocationDirectionsResponse(response DisplayTextsResponse) {
	r.listLocationDirectionsResponse = append(r.listLocationDirectionsResponse, response)
}
