package services

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/200sh/200sh-dashboard/internal/repository"
	"github.com/200sh/200sh-dashboard/models"
	"strings"
)

type UserService interface {
	GetBySubjectID(id string) (*models.User, error)
	Create(u *models.User) error
	Update(u *models.User) error
	//Deactivate(u *models.User) error
}
type userService struct {
	db   *sql.DB
	repo *repository.Queries
}

func (s *userService) GetBySubjectID(id string) (*models.User, error) {
	dbUser, err := s.repo.FindUserByProviderID(context.Background(), id)
	if err != nil {
		return nil, err
	}
	return models.FromDBUser(dbUser), nil
}

func NewUserService(db *sql.DB, repo *repository.Queries) UserService {
	return &userService{
		db,
		repo,
	}
}

func (s *userService) Update(u *models.User) error {
	if err := u.Validate(); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	return s.repo.UpdateUser(context.Background(), repository.UpdateUserParams{
		ID:    u.Id,
		Name:  strings.TrimSpace(u.Name),
		Email: strings.TrimSpace(u.Email),
	})
}

func (s *userService) Create(u *models.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	params := repository.CreateUserParams{
		ProviderID: u.ProviderId,
		Provider:   u.Provider,
		Name:       u.Name,
		Email:      u.Email,
		Status:     int64(u.Status),
	}

	result, err := s.repo.CreateUser(context.Background(), params)
	if err != nil {
		return err
	}

	u.Id = result.ID // Populate generated ID
	return nil
}
