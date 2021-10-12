package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

type NodePayload struct {
	Node  db.Node
	Error error
}

func (r *MockRepositoryService) CreateNode(ctx context.Context, arg db.CreateNodeParams) (db.Node, error) {
	if len(r.createNodePayload) == 0 {
		return db.Node{}, ErrorNotFound()
	}

	response := r.createNodePayload[0]
	r.createNodePayload = r.createNodePayload[1:]
	return response.Node, response.Error
}

func (r *MockRepositoryService) GetNode(ctx context.Context, id int64) (db.Node, error) {
	if len(r.getNodePayload) == 0 {
		return db.Node{}, ErrorNotFound()
	}

	response := r.getNodePayload[0]
	r.getNodePayload = r.getNodePayload[1:]
	return response.Node, response.Error
}

func (r *MockRepositoryService) SetCreateNodePayload(response NodePayload) {
	r.createNodePayload = append(r.createNodePayload, response)
}

func (r *MockRepositoryService) SetGetNodePayload(response NodePayload) {
	r.getNodePayload = append(r.getNodePayload, response)
}
