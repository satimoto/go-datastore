package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

type StatusSchedulesMockData struct {
	StatusSchedules []db.StatusSchedule
	Error           error
}
func (r *MockRepositoryService) CreateStatusSchedule(ctx context.Context, arg db.CreateStatusScheduleParams) (db.StatusSchedule, error) {
	r.createStatusScheduleMockData = append(r.createStatusScheduleMockData, arg)
	return db.StatusSchedule{}, nil
}

func (r *MockRepositoryService) ListStatusSchedules(ctx context.Context, evseID int64) ([]db.StatusSchedule, error) {
	if len(r.listStatusSchedulesMockData) == 0 {
		return []db.StatusSchedule{}, nil
	}

	response := r.listStatusSchedulesMockData[0]
	r.listStatusSchedulesMockData = r.listStatusSchedulesMockData[1:]
	return response.StatusSchedules, response.Error
}

func (r *MockRepositoryService) GetCreateStatusScheduleMockData() (db.CreateStatusScheduleParams, error) {
	if len(r.createStatusScheduleMockData) == 0 {
		return db.CreateStatusScheduleParams{}, ErrorNotFound()
	}

	response := r.createStatusScheduleMockData[0]
	r.createStatusScheduleMockData = r.createStatusScheduleMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetListStatusSchedulesMockData(response StatusSchedulesMockData) {
	r.listStatusSchedulesMockData = append(r.listStatusSchedulesMockData, response)
}
