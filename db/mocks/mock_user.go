package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

type UserResponse struct {
	User  db.User
	Error error
}

type UsersResponse struct {
	Users []db.User
	Error error
}

func (r *MockRepositoryService) CreateUser(ctx context.Context, arg db.CreateUserParams) (db.User, error) {
	return r.createUserResponse.User, r.createUserResponse.Error
}

func (r *MockRepositoryService) DeleteUser(ctx context.Context, id int64) error {
	return r.deleteUserResponse
}

func (r *MockRepositoryService) GetUser(ctx context.Context, id int64) (db.User, error) {
	return r.getUserResponse.User, r.getUserResponse.Error
}

func (r *MockRepositoryService) ListUsers(ctx context.Context) ([]db.User, error) {
	return r.listUsersResponse.Users, r.listUsersResponse.Error
}

func (r *MockRepositoryService) UpdateUser(ctx context.Context, arg db.UpdateUserParams) (db.User, error) {
	return r.updateUserResponse.User, r.updateUserResponse.Error
}

func (r *MockRepositoryService) SetCreateUserResponse(response UserResponse) {
	r.createUserResponse = response
}

func (r *MockRepositoryService) SetDeleteUserResponse(err error) {
	r.deleteUserResponse = err
}

func (r *MockRepositoryService) SetGetUserResponse(response UserResponse) {
	r.getUserResponse = response
}

func (r *MockRepositoryService) SetListUsersResponse(response UsersResponse) {
	r.listUsersResponse = response
}

func (r *MockRepositoryService) SetUpdateUserResponse(response UserResponse) {
	r.updateUserResponse = response
}
