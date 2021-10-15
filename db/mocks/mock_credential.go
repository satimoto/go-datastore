package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

type CredentialMockData struct {
	Credential db.Credential
	Error      error
}

func (r *MockRepositoryService) CreateCredential(ctx context.Context, arg db.CreateCredentialParams) (db.Credential, error) {
	r.createCredentialMockData = append(r.createCredentialMockData, arg)
	return db.Credential{}, nil
}

func (r *MockRepositoryService) GetCredentialByPartyAndCountryCode(ctx context.Context, arg db.GetCredentialByPartyAndCountryCodeParams) (db.Credential, error) {
	if len(r.getCredentialByPartyAndCountryCodeMockData) == 0 {
		return db.Credential{}, ErrorNotFound()
	}

	response := r.getCredentialByPartyAndCountryCodeMockData[0]
	r.getCredentialByPartyAndCountryCodeMockData = r.getCredentialByPartyAndCountryCodeMockData[1:]
	return response.Credential, response.Error
}

func (r *MockRepositoryService) GetCreateCredentialMockData() (db.CreateCredentialParams, error) {
	if len(r.createCredentialMockData) == 0 {
		return db.CreateCredentialParams{}, ErrorNotFound()
	}

	response := r.createCredentialMockData[0]
	r.createCredentialMockData = r.createCredentialMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetGetCredentialByPartyAndCountryCodeMockData(response CredentialMockData) {
	r.getCredentialByPartyAndCountryCodeMockData = append(r.getCredentialByPartyAndCountryCodeMockData, response)
}
