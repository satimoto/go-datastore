package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

func (r *MockRepositoryService) ListLocationDirections(ctx context.Context, locationID int64) ([]db.DisplayText, error) {
	return r.listLocationDirectionsResponse.DisplayTexts, r.listLocationDirectionsResponse.Error
}

func (r *MockRepositoryService) SetListLocationDirectionsResponse(response DisplayTextsResponse) {
	r.listLocationDirectionsResponse = response
}
