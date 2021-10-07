package db_test

import (
	"context"
	"testing"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-datastore/db/mocks"
)

func TestCreateNode(t *testing.T) {
	ctx := context.Background()
	mockRepositoryService := mocks.NewMockRepository()

	t.Run("Create Node", func(t *testing.T) {
		expectNode := db.Node{
			Pubkey:  "abcdef1234567890",
			Address: "192.168.0.1:1234",
		}

		mockRepositoryService.SetCreateNodeResponse(mocks.NodeResponse{Node: expectNode, Error: nil})

		node, err := mockRepositoryService.CreateNode(ctx, db.CreateNodeParams{
			Pubkey:  "abcdef1234567890",
			Address: "192.168.0.1:1234",
		})

		if node.Pubkey != expectNode.Pubkey {
			t.Errorf("Pubkey: got %v, want %v", expectNode.Pubkey, node.Pubkey)
		}

		if node.Address != expectNode.Address {
			t.Errorf("Address: got %v, want %v", expectNode.Address, node.Address)
		}

		if err != nil {
			t.Errorf("Error: got %v, want %v", err, nil)
		}
	})
}
