package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

type CredentialPayload struct {
	Credential db.Credential
	Error      error
}

func (r *MockRepositoryService) CreateCredential(ctx context.Context, arg db.CreateCredentialParams) (db.Credential, error) {
	if len(r.createCredentialPayload) == 0 {
		return db.Credential{}, ErrorNotFound()
	}

	response := r.createCredentialPayload[0]
	r.createCredentialPayload = r.createCredentialPayload[1:]
	return response.Credential, response.Error
}

func (r *MockRepositoryService) GetCredentialByPartyAndCountryCode(ctx context.Context, arg db.GetCredentialByPartyAndCountryCodeParams) (db.Credential, error) {
	if len(r.getCredentialByPartyAndCountryCodePayload) == 0 {
		return db.Credential{}, ErrorNotFound()
	}

	response := r.getCredentialByPartyAndCountryCodePayload[0]
	r.getCredentialByPartyAndCountryCodePayload = r.getCredentialByPartyAndCountryCodePayload[1:]
	return response.Credential, response.Error
}

func (r *MockRepositoryService) SetCreateCredentialPayload(response CredentialPayload) {
	r.createCredentialPayload = append(r.createCredentialPayload, response)
}

func (r *MockRepositoryService) SetGetCredentialByPartyAndCountryCodePayload(response CredentialPayload) {
	r.getCredentialByPartyAndCountryCodePayload = append(r.getCredentialByPartyAndCountryCodePayload, response)
}
