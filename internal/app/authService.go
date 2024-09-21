package app

import (
	"GoGateway/internal/domain"
	"GoGateway/internal/ports"
	"GoGateway/util"
	"GoGateway/util/errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService interface {
	GetUserByID(id string) (*domain.User, *errors.AppError)
	Authenticate(username, password string) (string, *errors.AppError)
	ValidateAndExtractClaims(token string) (jwt.MapClaims, bool, *errors.AppError)
}

type authService struct {
	repo        ports.AuthRepository
	httpClient  ports.HTTPClient
	externalAPI string
	logger      util.Logger
	jwtSecret   string // Secret key for signing JWT
}

func NewAuthService(repo ports.AuthRepository, httpClient ports.HTTPClient, externalAPI string, logger util.Logger, jwtSecret string) AuthService {
	return &authService{
		repo:        repo,
		httpClient:  httpClient,
		externalAPI: externalAPI,
		logger:      logger,
		jwtSecret:   jwtSecret,
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

// Authenticate generates a JWT token after validating the user's credentials
func (s *authService) Authenticate(username, password string) (string, *errors.AppError) {
	if username != "admin" || password != "password" {
		return "", errors.NewUnauthorizedError("Invalid credentials", nil)
	}

	// Create the JWT claims, which include the username and expiry time
	claims := &jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(), // 72 hours expiration
	}

	// Create the token with the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token using the secret key
	tokenString, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", errors.NewInternalServerError("Error generating token", err)
	}

	return tokenString, nil
}

// ValidateAndExtractClaims validates the JWT token and returns the claims
func (s *authService) ValidateAndExtractClaims(tokenString string) (jwt.MapClaims, bool, *errors.AppError) {
	if tokenString == "" {
		return nil, false, errors.NewAppError("Invalid token", http.StatusBadRequest)
	}

	// Parse the token and extract claims
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC (HS256)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.NewAppError("Unexpected signing method", http.StatusUnauthorized)
		}
		return []byte(s.jwtSecret), nil
	})

	if err != nil || !token.Valid {
		return nil, false, errors.NewAppError("Token validation failed", http.StatusUnauthorized)
	}

	// Extract claims from the token
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true, nil
	}

	return nil, false, errors.NewAppError("Invalid token claims", http.StatusUnauthorized)
}
