package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/pkg/db"
)

type RoutingEventMockData struct {
	RoutingEvent db.RoutingEvent
	Error        error
}

func (r *MockRepositoryService) CreateRoutingEvent(ctx context.Context, arg db.CreateRoutingEventParams) (db.RoutingEvent, error) {
	r.createRoutingEventMockData = append(r.createRoutingEventMockData, arg)
	return db.RoutingEvent{}, nil
}

func (r *MockRepositoryService) GetCreateRoutingEventMockData() (db.CreateRoutingEventParams, error) {
	if len(r.createRoutingEventMockData) == 0 {
		return db.CreateRoutingEventParams{}, ErrorNotFound()
	}

	response := r.createRoutingEventMockData[0]
	r.createRoutingEventMockData = r.createRoutingEventMockData[1:]
	return response, nil
}
