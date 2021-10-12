package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

type CredentialResponse struct {
	Credential db.Credential
	Error      error
}

func (r *MockRepositoryService) CreateCredential(ctx context.Context, arg db.CreateCredentialParams) (db.Credential, error) {
	if len(r.createCredentialResponse) == 0 {
		return db.Credential{}, ErrorNotFound()
	}

	response := r.createCredentialResponse[0]
	r.createCredentialResponse = r.createCredentialResponse[1:]
	return response.Credential, response.Error
}

func (r *MockRepositoryService) GetCredentialByPartyAndCountryCode(ctx context.Context, arg db.GetCredentialByPartyAndCountryCodeParams) (db.Credential, error) {
	if len(r.getCredentialByPartyAndCountryCodeResponse) == 0 {
		return db.Credential{}, ErrorNotFound()
	}

	response := r.getCredentialByPartyAndCountryCodeResponse[0]
	r.getCredentialByPartyAndCountryCodeResponse = r.getCredentialByPartyAndCountryCodeResponse[1:]
	return response.Credential, response.Error
}

func (r *MockRepositoryService) SetCreateCredentialResponse(response CredentialResponse) {
	r.createCredentialResponse = append(r.createCredentialResponse, response)
}

func (r *MockRepositoryService) SetGetCredentialByPartyAndCountryCodeResponse(response CredentialResponse) {
	r.getCredentialByPartyAndCountryCodeResponse = append(r.getCredentialByPartyAndCountryCodeResponse, response)
}
