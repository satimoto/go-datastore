package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

type UserMockData struct {
	User  db.User
	Error error
}

type UsersMockData struct {
	Users []db.User
	Error error
}

func (r *MockRepositoryService) CreateUser(ctx context.Context, arg db.CreateUserParams) (db.User, error) {
	r.createUserMockData = append(r.createUserMockData, arg)
	return db.User{}, nil
}

func (r *MockRepositoryService) DeleteUser(ctx context.Context, id int64) error {
	if len(r.deleteUserMockData) == 0 {
		return nil
	}

	response := r.deleteUserMockData[0]
	r.deleteUserMockData = r.deleteUserMockData[1:]
	return response
}

func (r *MockRepositoryService) GetUser(ctx context.Context, id int64) (db.User, error) {
	if len(r.getUserMockData) == 0 {
		return db.User{}, ErrorNotFound()
	}

	response := r.getUserMockData[0]
	r.getUserMockData = r.getUserMockData[1:]
	return response.User, response.Error
}

func (r *MockRepositoryService) ListUsers(ctx context.Context) ([]db.User, error) {
	if len(r.listUsersMockData) == 0 {
		return []db.User{}, nil
	}

	response := r.listUsersMockData[0]
	r.listUsersMockData = r.listUsersMockData[1:]
	return response.Users, response.Error
}

func (r *MockRepositoryService) UpdateUser(ctx context.Context, arg db.UpdateUserParams) (db.User, error) {
	if len(r.updateUserMockData) == 0 {
		return db.User{}, ErrorNotFound()
	}

	response := r.updateUserMockData[0]
	r.updateUserMockData = r.updateUserMockData[1:]
	return response.User, response.Error
}

func (r *MockRepositoryService) GetCreateUserMockData() (db.CreateUserParams, error) {
	if len(r.createUserMockData) == 0 {
		return db.CreateUserParams{}, ErrorNotFound()
	}

	response := r.createUserMockData[0]
	r.createUserMockData = r.createUserMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetDeleteUserMockData(err error) {
	r.deleteUserMockData = append(r.deleteUserMockData, err)
}

func (r *MockRepositoryService) SetGetUserMockData(response UserMockData) {
	r.getUserMockData = append(r.getUserMockData, response)
}

func (r *MockRepositoryService) SetListUsersMockData(response UsersMockData) {
	r.listUsersMockData = append(r.listUsersMockData, response)
}

func (r *MockRepositoryService) SetUpdateUserMockData(response UserMockData) {
	r.updateUserMockData = append(r.updateUserMockData, response)
}
