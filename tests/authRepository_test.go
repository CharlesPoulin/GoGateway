package db_test

import (
	"testing"

	"GoGateway/infra/db"
	"GoGateway/util"

	_ "github.com/lib/pq"
)

func TestGetUserByID(t *testing.T) {
	// Use a mock database or an in-memory test database
	connStr := "user=postgres dbname=testdb sslmode=disable"
	logger := util.NewLogger("info")
	repo, err := db.NewAuthRepository(connStr, logger)
	if err != nil {
		t.Fatalf("Failed to initialize repository: %v", err)
	}
	defer repo.Close()

	// Prepare mock data
	mockID := "123"
	_, err = repo.GetUserByID(mockID)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}
