package user

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"os"
	"sync"

	"github.com/sergey4qb/mf1-test/model"
)

var (
	fileRepoPath = "users.json"
)

type Repository interface {
	Create(ctx context.Context, user *model.User) error
	GetByID(ctx context.Context, id uuid.UUID) (*model.User, error)
	GetAll(ctx context.Context) ([]model.User, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type fileUserRepository struct {
	filePath string
	mu       sync.Mutex
}

func New() (Repository, error) {
	if err := initUserJsonFile(fileRepoPath); err != nil {
		return nil, err
	}
	return &fileUserRepository{filePath: fileRepoPath}, nil
}

func (r *fileUserRepository) Create(ctx context.Context, user *model.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	users, err := r.getAllNoLock()
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

func (r *fileUserRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	users, err := r.getAllNoLock()
	if err != nil {
		return nil, err
	}

	for _, u := range users {
		if u.ID == id {
			return &u, nil
		}
	}
	return nil, errUserNotFound
}

func (r *fileUserRepository) GetAll(ctx context.Context) ([]model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.getAllNoLock()
}

func (r *fileUserRepository) Update(ctx context.Context, user *model.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	users, err := r.getAllNoLock()
	if err != nil {
		return err
	}

	found := false
	for i, u := range users {
		if u.ID == user.ID {
			users[i] = *user
			found = true
			break
		}
	}
	if !found {
		return errUserNotFound
	}

	data, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(r.filePath, data, 0644); err != nil {
		return err
	}

	return nil
}

func (r *fileUserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	users, err := r.getAllNoLock()
	if err != nil {
		return err
	}

	index := -1
	for i, u := range users {
		if u.ID == id {
			index = i
			break
		}
	}
	if index == -1 {
		return errUserNotFound
	}

	users = append(users[:index], users[index+1:]...)

	data, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(r.filePath, data, 0644); err != nil {
		return err
	}

	return nil
}

func (r *fileUserRepository) getAllNoLock() ([]model.User, error) {
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
