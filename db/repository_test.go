package db_test

import (
	"context"
	"testing"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-datastore-mocks/db"
)

func TestCreateNode(t *testing.T) {
	ctx := context.Background()

	t.Run("Create Node", func(t *testing.T) {
		mockRepositoryService := mocks.NewMockRepositoryService()

		node := db.Node{
			Pubkey:  "abcdef1234567890",
			Address: "192.168.0.1:1234",
		}

		_, err := mockRepositoryService.CreateNode(ctx, db.CreateNodeParams{
			Pubkey:  "abcdef1234567890",
			Address: "192.168.0.1:1234",
		})

		mockNode, err := mockRepositoryService.GetCreateNodeMockData()

		if node.Pubkey != mockNode.Pubkey {
			t.Errorf("Pubkey: got %v, want %v", mockNode.Pubkey, node.Pubkey)
		}

		if node.Address != mockNode.Address {
			t.Errorf("Address: got %v, want %v", mockNode.Address, node.Address)
		}

		if err != nil {
			t.Errorf("Error: got %v, want %v", err, nil)
		}
	})
}
