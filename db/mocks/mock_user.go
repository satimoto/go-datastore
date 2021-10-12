package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

type UserPayload struct {
	User  db.User
	Error error
}

type UsersPayload struct {
	Users []db.User
	Error error
}

func (r *MockRepositoryService) CreateUser(ctx context.Context, arg db.CreateUserParams) (db.User, error) {
	if len(r.createUserPayload) == 0 {
		return db.User{}, ErrorNotFound()
	}

	response := r.createUserPayload[0]
	r.createUserPayload = r.createUserPayload[1:]
	return response.User, response.Error
}

func (r *MockRepositoryService) DeleteUser(ctx context.Context, id int64) error {
	return r.deleteUserPayload
}

func (r *MockRepositoryService) GetUser(ctx context.Context, id int64) (db.User, error) {
	if len(r.getUserPayload) == 0 {
		return db.User{}, ErrorNotFound()
	}

	response := r.getUserPayload[0]
	r.getUserPayload = r.getUserPayload[1:]
	return response.User, response.Error
}

func (r *MockRepositoryService) ListUsers(ctx context.Context) ([]db.User, error) {
	if len(r.listUsersPayload) == 0 {
		return []db.User{}, nil
	}

	response := r.listUsersPayload[0]
	r.listUsersPayload = r.listUsersPayload[1:]
	return response.Users, response.Error
}

func (r *MockRepositoryService) UpdateUser(ctx context.Context, arg db.UpdateUserParams) (db.User, error) {
	if len(r.updateUserPayload) == 0 {
		return db.User{}, ErrorNotFound()
	}

	response := r.updateUserPayload[0]
	r.updateUserPayload = r.updateUserPayload[1:]
	return response.User, response.Error
}

func (r *MockRepositoryService) SetCreateUserPayload(response UserPayload) {
	r.createUserPayload = append(r.createUserPayload, response)
}

func (r *MockRepositoryService) SetDeleteUserPayload(err error) {
	r.deleteUserPayload = err
}

func (r *MockRepositoryService) SetGetUserPayload(response UserPayload) {
	r.getUserPayload = append(r.getUserPayload, response)
}

func (r *MockRepositoryService) SetListUsersPayload(response UsersPayload) {
	r.listUsersPayload = append(r.listUsersPayload, response)
}

func (r *MockRepositoryService) SetUpdateUserPayload(response UserPayload) {
	r.updateUserPayload = append(r.updateUserPayload, response)
}
