package ports

import (
	"GoGateway/internal/domain"
	"GoGateway/util/errors"
)

type AuthRepository interface {
	GetUserByID(id string) (*domain.User, *errors.AppError)
	// Add the Close method to the interface
	Close() error
	// Define other repository methods as needed
}
