package db

import (
	"database/sql"
	"fmt"

	"GoGateway/internal/domain"
	"GoGateway/internal/ports"
	"GoGateway/util"
	"GoGateway/util/errors" // Import the correct package for AppError

	_ "github.com/lib/pq" // PostgreSQL driver
)

type AuthRepository struct {
	db     *sql.DB
	logger util.Logger
}

func NewAuthRepository(connStr string, logger util.Logger) (ports.AuthRepository, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Verify connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	logger.Info("Connected to the database")
	return &AuthRepository{db: db, logger: logger}, nil
}

// GetUserByID fetches a user by ID
func (r *AuthRepository) GetUserByID(id string) (*domain.User, *errors.AppError) {
	var user domain.User
	query := "SELECT id, username, email FROM users WHERE id = $1"
	row := r.db.QueryRow(query, id)
	if err := row.Scan(&user.ID, &user.Username, &user.Email); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NewNotFoundError("User not found", nil)
		}
		r.logger.Error("Database query failed", "error", err)
		return nil, errors.NewInternalServerError("Database error", err)
	}
	return &user, nil
}

// GetUsers fetches all users from the database
func (r *AuthRepository) GetUsers() ([]*domain.User, *errors.AppError) {
	query := "SELECT id, username, email FROM users"
	rows, err := r.db.Query(query)
	if err != nil {
		r.logger.Error("Database query failed", "error", err)
		return nil, errors.NewInternalServerError("Database error", err)
	}
	defer rows.Close()

	var users []*domain.User
	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email); err != nil {
			r.logger.Error("Failed to scan user row", "error", err)
			return nil, errors.NewInternalServerError("Failed to scan user row", err)
		}
		users = append(users, &user)
	}

	// Check for errors after iteration
	if err = rows.Err(); err != nil {
		r.logger.Error("Row iteration error", "error", err)
		return nil, errors.NewInternalServerError("Row iteration error", err)
	}

	return users, nil
}

func (r *AuthRepository) Close() error {
	return r.db.Close()
}
