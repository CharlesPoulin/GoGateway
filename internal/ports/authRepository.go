package ports

import (
	"GoGateway/internal/domain"
	"GoGateway/util/errors"
)

type AuthRepository interface {
	GetUserByID(id string) (*domain.User, *errors.AppError)
	GetUsers() ([]*domain.User, *errors.AppError)
	Close() error
	// Define other repository methods as needed
}
