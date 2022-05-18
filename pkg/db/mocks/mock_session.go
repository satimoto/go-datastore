package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/pkg/db"
)

type SessionMockData struct {
	Session db.Session
	Error   error
}

type SessionsMockData struct {
	Sessions []db.Session
	Error    error
}

func (r *MockRepositoryService) CreateSession(ctx context.Context, arg db.CreateSessionParams) (db.Session, error) {
	r.createSessionMockData = append(r.createSessionMockData, arg)
	return db.Session{}, nil
}

func (r *MockRepositoryService) GetSessionByAuthorizationID(ctx context.Context, authorizationID string) (db.Session, error) {
	if len(r.getSessionByAuthorizationIDMockData) == 0 {
		return db.Session{}, ErrorNotFound()
	}

	response := r.getSessionByAuthorizationIDMockData[0]
	r.getSessionByAuthorizationIDMockData = r.getSessionByAuthorizationIDMockData[1:]
	return response.Session, response.Error
}

func (r *MockRepositoryService) GetSessionByLastUpdated(ctx context.Context, arg db.GetSessionByLastUpdatedParams) (db.Session, error) {
	if len(r.getSessionByLastUpdatedMockData) == 0 {
		return db.Session{}, ErrorNotFound()
	}

	response := r.getSessionByLastUpdatedMockData[0]
	r.getSessionByLastUpdatedMockData = r.getSessionByLastUpdatedMockData[1:]
	return response.Session, response.Error
}

func (r *MockRepositoryService) GetSessionByUid(ctx context.Context, uid string) (db.Session, error) {
	if len(r.getSessionByUidMockData) == 0 {
		return db.Session{}, ErrorNotFound()
	}

	response := r.getSessionByUidMockData[0]
	r.getSessionByUidMockData = r.getSessionByUidMockData[1:]
	return response.Session, response.Error
}

func (r *MockRepositoryService) UpdateSessionByUid(ctx context.Context, arg db.UpdateSessionByUidParams) (db.Session, error) {
	r.updateSessionByUidMockData = append(r.updateSessionByUidMockData, arg)
	return db.Session{}, nil
}

func (r *MockRepositoryService) GetCreateSessionMockData() (db.CreateSessionParams, error) {
	if len(r.createSessionMockData) == 0 {
		return db.CreateSessionParams{}, ErrorNotFound()
	}

	response := r.createSessionMockData[0]
	r.createSessionMockData = r.createSessionMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetUpdateSessionByUidMockData() (db.UpdateSessionByUidParams, error) {
	if len(r.updateSessionByUidMockData) == 0 {
		return db.UpdateSessionByUidParams{}, ErrorNotFound()
	}

	response := r.updateSessionByUidMockData[0]
	r.updateSessionByUidMockData = r.updateSessionByUidMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetGetSessionByAuthorizationIDMockData(response SessionMockData) {
	r.getSessionByAuthorizationIDMockData = append(r.getSessionByAuthorizationIDMockData, response)
}

func (r *MockRepositoryService) SetGetSessionByLastUpdatedMockData(response SessionMockData) {
	r.getSessionByLastUpdatedMockData = append(r.getSessionByLastUpdatedMockData, response)
}

func (r *MockRepositoryService) SetGetSessionByUidMockData(response SessionMockData) {
	r.getSessionByUidMockData = append(r.getSessionByUidMockData, response)
}
