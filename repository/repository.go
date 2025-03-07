package repository

import "github.com/sergey4qb/mf1-test/repository/user"

type Repository interface {
	GetUser() user.Repository
}

type repository struct {
	user user.Repository
}

func New() (Repository, error) {
	user, err := user.New()
	if err != nil {
		return nil, err
	}

	return &repository{
		user: user,
	}, nil
}

func (r *repository) GetUser() user.Repository {
	return r.user
}
