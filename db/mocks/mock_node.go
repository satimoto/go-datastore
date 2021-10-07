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
	return r.createNodeResponse.Node, r.createNodeResponse.Error
}

func (r *MockRepositoryService) GetNode(ctx context.Context, id int64) (db.Node, error) {
	return r.getNodeResponse.Node, r.getNodeResponse.Error
}

func (r *MockRepositoryService) SetCreateNodeResponse(response NodeResponse) {
	r.createNodeResponse = response
}

func (r *MockRepositoryService) SetGetNodeResponse(response NodeResponse) {
	r.getNodeResponse = response
}
