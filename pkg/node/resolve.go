package node

import (
	"context"

	"github.com/satimoto/go-datastore/pkg/db"
)

type NodeRepository interface {
	CreateNode(ctx context.Context, arg db.CreateNodeParams) (db.Node, error)
	GetNode(ctx context.Context, id int64) (db.Node, error)
	GetNodeByPubkey(ctx context.Context, pubkey string) (db.Node, error)
	GetNodeByUserID(ctx context.Context, id int64) (db.Node, error)
	ListNodes(ctx context.Context) ([]db.Node, error)
	UpdateNode(ctx context.Context, arg db.UpdateNodeParams) (db.Node, error)
}

type NodeResolver struct {
	Repository NodeRepository
}

func NewResolver(repositoryService *db.RepositoryService) *NodeResolver {
	repo := NodeRepository(repositoryService)

	return &NodeResolver{
		Repository: repo,
	}
}
