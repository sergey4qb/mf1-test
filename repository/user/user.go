package user

import (
	"encoding/json"
	"os"

	"github.com/sergey4qb/mf1-test/model"
)

var (
	fileRepoPath = "users.json"
)

type Repository interface {
	Create(user *model.User) error
	GetAll() ([]model.User, error)
}

type fileUserRepository struct {
	filePath string
}

func New() (Repository, error) {
	if err := initUserJsonFile(fileRepoPath); err != nil {
		return nil, err
	}
	return &fileUserRepository{filePath: fileRepoPath}, nil
}

func (r *fileUserRepository) Create(user *model.User) error {
	users, err := r.GetAll()
	if err != nil {
		return err
	}

	users = append(users, *user)

	data, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(r.filePath, data, 0644); err != nil {
		return err
	}

	return nil
}

func (r *fileUserRepository) GetAll() ([]model.User, error) {
	if _, err := os.Stat(r.filePath); os.IsNotExist(err) {
		return []model.User{}, nil
	}

	data, err := os.ReadFile(r.filePath)
	if err != nil {
		return nil, err
	}

	var users []model.User
	if err := json.Unmarshal(data, &users); err != nil {
		return nil, err
	}

	return users, nil
}
