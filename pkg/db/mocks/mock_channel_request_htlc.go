package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/pkg/db"
)

type ChannelRequestHtlcMockData struct {
	ChannelRequestHtlc db.ChannelRequestHtlc
	Error              error
}

type ChannelRequestHtlcsMockData struct {
	ChannelRequestHtlcs []db.ChannelRequestHtlc
	Error               error
}

func (r *MockRepositoryService) CreateChannelRequestHtlc(ctx context.Context, arg db.CreateChannelRequestHtlcParams) (db.ChannelRequestHtlc, error) {
	r.createChannelRequestHtlcMockData = append(r.createChannelRequestHtlcMockData, arg)
	return db.ChannelRequestHtlc{}, nil
}

func (r *MockRepositoryService) GetChannelRequestHtlc(ctx context.Context, channelRequestID int64) (db.ChannelRequestHtlc, error) {
	if len(r.getChannelRequestHtlcMockData) == 0 {
		return db.ChannelRequestHtlc{}, ErrorNotFound()
	}

	response := r.getChannelRequestHtlcMockData[0]
	r.getChannelRequestHtlcMockData = r.getChannelRequestHtlcMockData[1:]
	return response.ChannelRequestHtlc, response.Error
}

func (r *MockRepositoryService) GetChannelRequestHtlcByCircuitKey(ctx context.Context, arg db.GetChannelRequestHtlcByCircuitKeyParams) (db.ChannelRequestHtlc, error) {
	if len(r.getChannelRequestHtlcByCircuitKeyMockData) == 0 {
		return db.ChannelRequestHtlc{}, ErrorNotFound()
	}

	response := r.getChannelRequestHtlcByCircuitKeyMockData[0]
	r.getChannelRequestHtlcByCircuitKeyMockData = r.getChannelRequestHtlcByCircuitKeyMockData[1:]
	return response.ChannelRequestHtlc, response.Error
}

func (r *MockRepositoryService) ListChannelRequestHtlcs(ctx context.Context, channelRequestID int64) ([]db.ChannelRequestHtlc, error) {
	if len(r.listChannelRequestHtlcsMockData) == 0 {
		return []db.ChannelRequestHtlc{}, nil
	}

	response := r.listChannelRequestHtlcsMockData[0]
	r.listChannelRequestHtlcsMockData = r.listChannelRequestHtlcsMockData[1:]
	return response.ChannelRequestHtlcs, response.Error
}

func (r *MockRepositoryService) ListUnsettledChannelRequestHtlcs(ctx context.Context, channelRequestID int64) ([]db.ChannelRequestHtlc, error) {
	if len(r.listUnsettledChannelRequestHtlcsMockData) == 0 {
		return []db.ChannelRequestHtlc{}, nil
	}

	response := r.listUnsettledChannelRequestHtlcsMockData[0]
	r.listUnsettledChannelRequestHtlcsMockData = r.listUnsettledChannelRequestHtlcsMockData[1:]
	return response.ChannelRequestHtlcs, response.Error
}

func (r *MockRepositoryService) UpdateChannelRequestHtlc(ctx context.Context, arg db.UpdateChannelRequestHtlcParams) (db.ChannelRequestHtlc, error) {
	r.updateChannelRequestHtlcMockData = append(r.updateChannelRequestHtlcMockData, arg)
	return db.ChannelRequestHtlc{}, nil
}

func (r *MockRepositoryService) GetCreateChannelRequestHtlcMockData() (db.CreateChannelRequestHtlcParams, error) {
	if len(r.createChannelRequestHtlcMockData) == 0 {
		return db.CreateChannelRequestHtlcParams{}, ErrorNotFound()
	}

	response := r.createChannelRequestHtlcMockData[0]
	r.createChannelRequestHtlcMockData = r.createChannelRequestHtlcMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetGetChannelRequestHtlcMockData(response ChannelRequestHtlcMockData) {
	r.getChannelRequestHtlcMockData = append(r.getChannelRequestHtlcMockData, response)
}

func (r *MockRepositoryService) SetGetChannelRequestHtlcByCircuitKeyMockData(response ChannelRequestHtlcMockData) {
	r.getChannelRequestHtlcByCircuitKeyMockData = append(r.getChannelRequestHtlcByCircuitKeyMockData, response)
}

func (r *MockRepositoryService) SetListChannelRequestHtlcsMockData(response ChannelRequestHtlcsMockData) {
	r.listChannelRequestHtlcsMockData = append(r.listChannelRequestHtlcsMockData, response)
}

func (r *MockRepositoryService) SetListUnsettledChannelRequestHtlcsMockData(response ChannelRequestHtlcsMockData) {
	r.listUnsettledChannelRequestHtlcsMockData = append(r.listUnsettledChannelRequestHtlcsMockData, response)
}

func (r *MockRepositoryService) GetUpdateChannelRequestHtlcMockData() (db.UpdateChannelRequestHtlcParams, error) {
	if len(r.updateChannelRequestHtlcMockData) == 0 {
		return db.UpdateChannelRequestHtlcParams{}, ErrorNotFound()
	}

	response := r.updateChannelRequestHtlcMockData[0]
	r.updateChannelRequestHtlcMockData = r.updateChannelRequestHtlcMockData[1:]
	return response, nil
}
