package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/pkg/db"
)

type CountryAccountMockData struct {
	CountryAccount db.CountryAccount
	Error          error
}

type CountryAccountsMockData struct {
	CountryAccounts []db.CountryAccount
	Error           error
}

func (r *MockRepositoryService) CreateCountryAccount(ctx context.Context, arg db.CreateCountryAccountParams) (db.CountryAccount, error) {
	r.createCountryAccountMockData = append(r.createCountryAccountMockData, arg)
	return db.CountryAccount{}, nil
}

func (r *MockRepositoryService) GetCountryAccountByCountry(ctx context.Context, country string) (db.CountryAccount, error) {
	if len(r.getCountryAccountByCountryMockData) == 0 {
		return db.CountryAccount{}, ErrorNotFound()
	}

	response := r.getCountryAccountByCountryMockData[0]
	r.getCountryAccountByCountryMockData = r.getCountryAccountByCountryMockData[1:]
	return response.CountryAccount, response.Error
}

func (r *MockRepositoryService) ListCountryAccounts(ctx context.Context) ([]db.CountryAccount, error) {
	if len(r.listCountryAccountsMockData) == 0 {
		return []db.CountryAccount{}, nil
	}

	response := r.listCountryAccountsMockData[0]
	r.listCountryAccountsMockData = r.listCountryAccountsMockData[1:]
	return response.CountryAccounts, response.Error
}

func (r *MockRepositoryService) GetCreateCountryAccountMockData() (db.CreateCountryAccountParams, error) {
	if len(r.createCountryAccountMockData) == 0 {
		return db.CreateCountryAccountParams{}, ErrorNotFound()
	}

	response := r.createCountryAccountMockData[0]
	r.createCountryAccountMockData = r.createCountryAccountMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetGetCountryAccountByCountryMockData(response CountryAccountMockData) {
	r.getCountryAccountByCountryMockData = append(r.getCountryAccountByCountryMockData, response)
}

func (r *MockRepositoryService) SetListCountryAccountsMockData(response CountryAccountsMockData) {
	r.listCountryAccountsMockData = append(r.listCountryAccountsMockData, response)
}
