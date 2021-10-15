package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

type NodeMockData struct {
	Node  db.Node
	Error error
}

func (r *MockRepositoryService) CreateNode(ctx context.Context, arg db.CreateNodeParams) (db.Node, error) {
	r.createNodeMockData = append(r.createNodeMockData, arg)
	return db.Node{}, nil
}

func (r *MockRepositoryService) GetNode(ctx context.Context, id int64) (db.Node, error) {
	if len(r.getNodeMockData) == 0 {
		return db.Node{}, ErrorNotFound()
	}

	response := r.getNodeMockData[0]
	r.getNodeMockData = r.getNodeMockData[1:]
	return response.Node, response.Error
}

func (r *MockRepositoryService) GetCreateNodeMockData() (db.CreateNodeParams, error) {
	if len(r.createNodeMockData) == 0 {
		return db.CreateNodeParams{}, ErrorNotFound()
	}

	response := r.createNodeMockData[0]
	r.createNodeMockData = r.createNodeMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetGetNodeMockData(response NodeMockData) {
	r.getNodeMockData = append(r.getNodeMockData, response)
}
