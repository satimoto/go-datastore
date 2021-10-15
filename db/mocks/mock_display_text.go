package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

type DisplayTextsMockData struct {
	DisplayTexts []db.DisplayText
	Error        error
}

func (r *MockRepositoryService) CreateDisplayText(ctx context.Context, arg db.CreateDisplayTextParams) (db.DisplayText, error) {
	r.createDisplayTextMockData = append(r.createDisplayTextMockData, arg)
	return db.DisplayText{}, nil
}

func (r *MockRepositoryService) GetCreateDisplayTextMockData() (db.CreateDisplayTextParams, error) {
	if len(r.createDisplayTextMockData) == 0 {
		return db.CreateDisplayTextParams{}, ErrorNotFound()
	}

	response := r.createDisplayTextMockData[0]
	r.createDisplayTextMockData = r.createDisplayTextMockData[1:]
	return response, nil
}