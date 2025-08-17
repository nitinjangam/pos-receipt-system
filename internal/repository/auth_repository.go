package repository

import (
	"context"
	"database/sql"

	v1 "github.com/nitinjangam/pos-receipt-system/api/v1"
)

// AuthRepositoryInterface defines the methods for the auth repository.
type AuthRepositoryInterface interface {
	GetUserByUsername(ctx context.Context, username string) (v1.PostAuthLoginJSONBody, error)
	CreateUser(ctx context.Context, user v1.PostAuthLoginJSONBody) error
}

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (r *AuthRepository) GetUserByUsername(ctx context.Context, username string) (v1.PostAuthLoginJSONBody, error) {
	var user v1.PostAuthLoginJSONBody

	query := "SELECT username, password FROM auth_users WHERE username = ?"
	err := r.db.QueryRowContext(ctx, query, username).Scan(&user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, nil // User not found
		}
		return user, err // Other error
	}

	return user, nil
}

func (r *AuthRepository) CreateUser(ctx context.Context, user v1.PostAuthLoginJSONBody) error {
	query := "INSERT INTO auth_users (username, password) VALUES (?, ?)"
	_, err := r.db.ExecContext(ctx, query, user.Username, user.Password)
	if err != nil {
		return err // Return error if insertion fails
	}
	return nil // Return nil if insertion is successful
}
