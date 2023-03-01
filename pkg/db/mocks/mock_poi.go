package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/pkg/db"
)

type PoiMockData struct {
	Poi   db.Poi
	Error error
}

type PoisMockData struct {
	Pois  []db.Poi
	Error error
}

func (r *MockRepositoryService) CreatePoi(ctx context.Context, arg db.CreatePoiParams) (db.Poi, error) {
	r.createPoiMockData = append(r.createPoiMockData, arg)
	return db.Poi{}, nil
}

func (r *MockRepositoryService) DeletePoiByUid(ctx context.Context, uid string) error {
	r.deletePoiByUidMockData = append(r.deletePoiByUidMockData, uid)
	return nil
}

func (r *MockRepositoryService) GetPoi(ctx context.Context, id int64) (db.Poi, error) {
	if len(r.getPoiMockData) == 0 {
		return db.Poi{}, ErrorNotFound()
	}

	response := r.getPoiMockData[0]
	r.getPoiMockData = r.getPoiMockData[1:]
	return response.Poi, response.Error
}

func (r *MockRepositoryService) GetPoiByLastUpdated(ctx context.Context) (db.Poi, error) {
	if len(r.getPoiByLastUpdatedMockData) == 0 {
		return db.Poi{}, ErrorNotFound()
	}

	response := r.getPoiByLastUpdatedMockData[0]
	r.getPoiByLastUpdatedMockData = r.getPoiByLastUpdatedMockData[1:]
	return response.Poi, response.Error
}

func (r *MockRepositoryService) ListPoisByGeom(ctx context.Context, arg db.ListPoisByGeomParams) ([]db.Poi, error) {
	if len(r.listPoisByGeomMockData) == 0 {
		return []db.Poi{}, nil
	}

	response := r.listPoisByGeomMockData[0]
	r.listPoisByGeomMockData = r.listPoisByGeomMockData[1:]
	return response.Pois, response.Error
}

func (r *MockRepositoryService) UpdatePoiByUid(ctx context.Context, arg db.UpdatePoiByUidParams) (db.Poi, error) {
	r.updatePoiByUidMockData = append(r.updatePoiByUidMockData, arg)
	return db.Poi{}, nil
}

func (r *MockRepositoryService) GetCreatePoiMockData() (db.CreatePoiParams, error) {
	if len(r.createPoiMockData) == 0 {
		return db.CreatePoiParams{}, ErrorNotFound()
	}

	response := r.createPoiMockData[0]
	r.createPoiMockData = r.createPoiMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetDeletePoiByUidMockData() (string, error) {
	if len(r.deletePoiByUidMockData) == 0 {
		return "", ErrorNotFound()
	}

	response := r.deletePoiByUidMockData[0]
	r.deletePoiByUidMockData = r.deletePoiByUidMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetUpdatePoiByUidMockData() (db.UpdatePoiByUidParams, error) {
	if len(r.updatePoiByUidMockData) == 0 {
		return db.UpdatePoiByUidParams{}, ErrorNotFound()
	}

	response := r.updatePoiByUidMockData[0]
	r.updatePoiByUidMockData = r.updatePoiByUidMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetGetPoiMockData(response PoiMockData) {
	r.getPoiMockData = append(r.getPoiMockData, response)
}

func (r *MockRepositoryService) SetGetPoiByLastUpdatedMockData(response PoiMockData) {
	r.getPoiByLastUpdatedMockData = append(r.getPoiByLastUpdatedMockData, response)
}

func (r *MockRepositoryService) SetListPoisByGeomMockData(response PoisMockData) {
	r.listPoisByGeomMockData = append(r.listPoisByGeomMockData, response)
}
