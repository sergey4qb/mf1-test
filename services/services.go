package services

import (
	"github.com/sergey4qb/mf1-test/repository"
	"github.com/sergey4qb/mf1-test/services/user"
)

type Services interface {
	GetUser() user.User
}

type services struct {
	user user.User
}

func New(repository repository.Repository) (Services, error) {
	return &services{
		user: user.New(repository.GetUser()),
	}, nil
}

func (r *services) GetUser() user.User {
	return r.user
}
