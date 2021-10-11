package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

type NodeResponse struct {
	Node  db.Node
	Error error
}

func (r *MockRepositoryService) CreateNode(ctx context.Context, arg db.CreateNodeParams) (db.Node, error) {
	if len(r.createNodeResponse) == 0 {
		return db.Node{}, ErrorNotFound()
	}

	response := r.createNodeResponse[0]
	r.createNodeResponse = r.createNodeResponse[1:]
	return response.Node, response.Error
}

func (r *MockRepositoryService) GetNode(ctx context.Context, id int64) (db.Node, error) {
	if len(r.getNodeResponse) == 0 {
		return db.Node{}, ErrorNotFound()
	}

	response := r.getNodeResponse[0]
	r.getNodeResponse = r.getNodeResponse[1:]
	return response.Node, response.Error
}

func (r *MockRepositoryService) SetCreateNodeResponse(response NodeResponse) {
	r.createNodeResponse = append(r.createNodeResponse, response)
}

func (r *MockRepositoryService) SetGetNodeResponse(response NodeResponse) {
	r.getNodeResponse = append(r.getNodeResponse, response)
}
