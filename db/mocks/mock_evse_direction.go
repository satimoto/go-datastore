package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

func (r *MockRepositoryService) ListEvseDirections(ctx context.Context, evseID int64) ([]db.DisplayText, error) {
	return r.listEvseDirectionsResponse.DisplayTexts, r.listEvseDirectionsResponse.Error
}

func (r *MockRepositoryService) SetListEvseDirectionsResponse(response DisplayTextsResponse) {
	r.listEvseDirectionsResponse = response
}
