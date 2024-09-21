package app

import (
	"GoGateway/internal/domain"
	"GoGateway/internal/ports"
	"GoGateway/util"
	"GoGateway/util/errors"
)

type AuthService interface {
	GetUserByID(id string) (*domain.User, *errors.AppError)
	Authenticate(username, password string) (string, *errors.AppError)
	ValidateToken(token string) (bool, *errors.AppError) // This method must be implemented
}

type authService struct {
	repo        ports.AuthRepository
	httpClient  ports.HTTPClient
	externalAPI string
	logger      util.Logger
}

func NewAuthService(repo ports.AuthRepository, httpClient ports.HTTPClient, externalAPI string, logger util.Logger) AuthService {
	return &authService{
		repo:        repo,
		httpClient:  httpClient,
		externalAPI: externalAPI,
		logger:      logger,
	}
}

// GetUserByID fetches a user by ID
func (s *authService) GetUserByID(id string) (*domain.User, *errors.AppError) {
	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return nil, errors.NewNotFoundError("User not found", err)
	}
	return user, nil
}

// Authenticate authenticates the user with username and password
func (s *authService) Authenticate(username, password string) (string, *errors.AppError) {
	if username != "admin" || password != "password" {
		return "", errors.NewUnauthorizedError("Invalid credentials", nil)
	}
	token := "some-generated-token" // Replace with actual token generation logic
	return token, nil
}

// ValidateToken checks the validity of the given token
func (s *authService) ValidateToken(token string) (bool, *errors.AppError) {
	// Placeholder logic: Implement actual token validation
	if token == "valid-token" {
		return true, nil
	}
	return false, errors.NewUnauthorizedError("Invalid token", nil)
}
