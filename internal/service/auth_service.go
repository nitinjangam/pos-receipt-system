package service

import (
	"context"
	"errors"
	"fmt"

	v1 "github.com/nitinjangam/pos-receipt-system/api/v1"
	"github.com/nitinjangam/pos-receipt-system/internal/repository"
	"go.uber.org/zap"
)

var emptyUser = v1.PostAuthLoginJSONBody{}

type AuthServiceInterface interface {
	Login(ctx context.Context, username string, password string) error
	Register(ctx context.Context, username string, password string) error
}

type AuthService struct {
	authRepo *repository.AuthRepository
	logger   *zap.SugaredLogger
}

func NewAuthService(authRepository *repository.AuthRepository, logger *zap.SugaredLogger) AuthServiceInterface {
	return &AuthService{
		authRepo: authRepository,
		logger:   logger,
	}
}

func (s *AuthService) Login(ctx context.Context, username string, password string) error {
	// Get user from repository
	user, err := s.authRepo.GetUserByUsername(ctx, username)
	if err != nil {
		s.logger.Errorw("Failed to get user by username", "username", username, "error", err)
		return err
	}
	if user == emptyUser {
		return errors.New("user not found")
	}
	// simple password check
	if *user.Password != password {
		return errors.New("wrong password")
	}
	return nil
}

func (s *AuthService) Register(ctx context.Context, username string, password string) error {
	// check if user already exists
	user, err := s.authRepo.GetUserByUsername(ctx, username)
	if err != nil {
		s.logger.Errorw("Failed to get user by username", "username", username, "error", err)
		return err
	}
	if user.Username != nil {
		return errors.New(fmt.Sprintf("username %s is already taken", username))
	}
	// store the user in the repository
	newUser := v1.PostAuthLoginJSONBody{
		Username: &username,
		Password: &password,
	}
	err = s.authRepo.CreateUser(ctx, newUser)
	if err != nil {
		return err
	}
	s.logger.Infow("User registered successfully", "username", username)
	return nil
}
