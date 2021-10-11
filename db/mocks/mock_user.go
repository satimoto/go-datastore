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
	if len(r.createUserResponse) == 0 {
		return db.User{}, ErrorNotFound()
	}

	response := r.createUserResponse[0]
	r.createUserResponse = r.createUserResponse[1:]
	return response.User, response.Error
}

func (r *MockRepositoryService) DeleteUser(ctx context.Context, id int64) error {
	return r.deleteUserResponse
}

func (r *MockRepositoryService) GetUser(ctx context.Context, id int64) (db.User, error) {
	if len(r.getUserResponse) == 0 {
		return db.User{}, ErrorNotFound()
	}

	response := r.getUserResponse[0]
	r.getUserResponse = r.getUserResponse[1:]
	return response.User, response.Error
}

func (r *MockRepositoryService) ListUsers(ctx context.Context) ([]db.User, error) {
	if len(r.listUsersResponse) == 0 {
		return []db.User{}, nil
	}

	response := r.listUsersResponse[0]
	r.listUsersResponse = r.listUsersResponse[1:]
	return response.Users, response.Error
}

func (r *MockRepositoryService) UpdateUser(ctx context.Context, arg db.UpdateUserParams) (db.User, error) {
	if len(r.updateUserResponse) == 0 {
		return db.User{}, ErrorNotFound()
	}

	response := r.updateUserResponse[0]
	r.updateUserResponse = r.updateUserResponse[1:]
	return response.User, response.Error
}

func (r *MockRepositoryService) SetCreateUserResponse(response UserResponse) {
	r.createUserResponse = append(r.createUserResponse, response)
}

func (r *MockRepositoryService) SetDeleteUserResponse(err error) {
	r.deleteUserResponse = err
}

func (r *MockRepositoryService) SetGetUserResponse(response UserResponse) {
	r.getUserResponse = append(r.getUserResponse, response)
}

func (r *MockRepositoryService) SetListUsersResponse(response UsersResponse) {
	r.listUsersResponse = append(r.listUsersResponse, response)
}

func (r *MockRepositoryService) SetUpdateUserResponse(response UserResponse) {
	r.updateUserResponse = append(r.updateUserResponse, response)
}
