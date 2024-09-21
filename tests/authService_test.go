package app_test

import (
	"testing"

	"GoGateway/internal/app"
	"GoGateway/internal/domain"
	"GoGateway/util"
	"GoGateway/util/errors"
)

// MockAuthRepository is a mock implementation of the AuthRepository interface
type MockAuthRepository struct {
	user *domain.User
	err  *errors.AppError
}

func (m *MockAuthRepository) GetUserByID(id string) (*domain.User, *errors.AppError) {
	return m.user, m.err
}

func TestGetUserByID_Success(t *testing.T) {
	logger := util.NewLogger("info")
	repo := &MockAuthRepository{
		user: &domain.User{ID: "123", Username: "testuser", Email: "test@example.com"},
		err:  nil,
	}
	authService := app.NewAuthService(repo, nil, "", logger)

	user, err := authService.GetUserByID("123")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if user.Username != "testuser" {
		t.Errorf("Expected username to be 'testuser', got %v", user.Username)
	}
}

func TestGetUserByID_NotFound(t *testing.T) {
	logger := util.NewLogger("info")
	repo := &MockAuthRepository{
		user: nil,
		err:  errors.NewNotFoundError("User not found", nil),
	}
	authService := app.NewAuthService(repo, nil, "", logger)

	_, err := authService.GetUserByID("123")
	if err == nil || err.GetStatusCode() != 404 {
		t.Errorf("Expected not found error, got %v", err)
	}
}
