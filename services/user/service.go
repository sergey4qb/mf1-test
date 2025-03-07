package user

import (
	"context"
	"github.com/google/uuid"
	"github.com/sergey4qb/mf1-test/dto"
	"github.com/sergey4qb/mf1-test/model"
	"github.com/sergey4qb/mf1-test/repository/user"
)

type User interface {
	Create(ctx context.Context, user *model.User) error
	GetByID(ctx context.Context, id uuid.UUID) (*model.User, error)
	GetAll(ctx context.Context) ([]model.User, error)
	Update(ctx context.Context, dto *dto.UpdateUserDTO) (*model.User, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type service struct {
	repo user.Repository
}

func New(repo user.Repository) User {
	return &service{repo: repo}
}

func (s *service) Create(ctx context.Context, user *model.User) error {
	if user.Name == "" {
		return errInvalidName
	}

	if user.Email == "" {
		return errInvalidEmail
	}

	if !emailRegex.MatchString(user.Email) {
		return errInvalidFormatEmail
	}

	user.ID = uuid.New()

	return s.repo.Create(ctx, user)
}

func (s *service) GetByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *service) GetAll(ctx context.Context) ([]model.User, error) {
	return s.repo.GetAll(ctx)
}

func (s *service) Update(ctx context.Context, dto *dto.UpdateUserDTO) (*model.User, error) {
	existingUser, err := s.GetByID(ctx, dto.ID)
	if err != nil {
		return nil, err
	}
	if dto.Name != nil {
		if *dto.Name == "" {
			return nil, errInvalidName
		}
		existingUser.Name = *dto.Name
	}
	if dto.Email != nil {
		if *dto.Email == "" {
			return nil, errInvalidEmail
		}
		if !emailRegex.MatchString(*dto.Email) {
			return nil, errInvalidFormatEmail
		}
		existingUser.Email = *dto.Email
	}
	if err := s.repo.Update(ctx, existingUser); err != nil {
		return nil, err
	}
	return existingUser, nil
}

func (s *service) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}
