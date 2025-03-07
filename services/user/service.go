package user

import (
	"github.com/google/uuid"
	"github.com/sergey4qb/mf1-test/model"
	"github.com/sergey4qb/mf1-test/repository/user"
)

type User interface {
	Create(user *model.User) error
	GetAll() ([]model.User, error)
}

type service struct {
	repo user.Repository
}

func New(repo user.Repository) User {
	return &service{repo: repo}
}

func (s *service) Create(user *model.User) error {
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

	return s.repo.Create(user)
}

func (s *service) GetAll() ([]model.User, error) {
	return s.repo.GetAll()
}
