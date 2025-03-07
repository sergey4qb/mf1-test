package user

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/uuid"
	"github.com/sergey4qb/mf1-test/model"
	"github.com/stretchr/testify/assert"
)

func TestFileUserRepository_CreateAndGetAll(t *testing.T) {
	tempDir := t.TempDir()
	fileRepoPath = filepath.Join(tempDir, "users_test.json")

	repo, err := New()
	assert.NoError(t, err)

	users, err := repo.GetAll()
	assert.NoError(t, err)
	assert.Len(t, users, 0)

	newUser := &model.User{
		ID:    uuid.New(),
		Name:  "Test User",
		Email: "test@example.com",
	}
	err = repo.Create(newUser)
	assert.NoError(t, err)

	users, err = repo.GetAll()
	assert.NoError(t, err)
	assert.Len(t, users, 1)
	assert.Equal(t, newUser.Name, users[0].Name)
	assert.Equal(t, newUser.Email, users[0].Email)

	_ = os.Remove(fileRepoPath)

}

func TestFileUserRepository_GetAll_FileNotExist(t *testing.T) {
	tempDir := t.TempDir()
	fileRepoPath = filepath.Join(tempDir, "nonexistent.json")
	_ = os.Remove(fileRepoPath)

	repo, err := New()
	assert.NoError(t, err)

	users, err := repo.GetAll()
	assert.NoError(t, err)
	assert.Len(t, users, 0)

}

func TestFileUserRepository_GetAll_InvalidJSON(t *testing.T) {
	tempDir := t.TempDir()
	fileRepoPath = filepath.Join(tempDir, "users.json")

	err := os.WriteFile(fileRepoPath, []byte("invalid json"), 0644)
	assert.NoError(t, err)

	repo, err := New()
	assert.NoError(t, err)

	users, err := repo.GetAll()
	assert.Error(t, err, "expected error on invalid json")
	assert.Nil(t, users, "expected nil users on invalid json")

}

func TestFileUserRepository_Create_ErrorOnWrite(t *testing.T) {
	tempDir := t.TempDir()
	fileRepoPath = filepath.Join(tempDir, "users.json")

	err := os.WriteFile(fileRepoPath, []byte("[]"), 0444)
	assert.NoError(t, err)

	repo, err := New()
	assert.NoError(t, err)

	newUser := &model.User{
		ID:    uuid.New(),
		Name:  "Test User",
		Email: "test@example.com",
	}
	err = repo.Create(newUser)
	assert.Error(t, err)
}

func TestFileUserRepository_ContentPersistence(t *testing.T) {
	tempDir := t.TempDir()
	fileRepoPath = filepath.Join(tempDir, "users.json")
	_ = os.Remove(fileRepoPath)

	repo, err := New()
	assert.NoError(t, err)

	newUser := &model.User{
		ID:    uuid.New(),
		Name:  "Persistent User",
		Email: "persist@example.com",
	}
	err = repo.Create(newUser)
	assert.NoError(t, err)

	data, err := os.ReadFile(fileRepoPath)
	assert.NoError(t, err)

	var users []model.User
	err = json.Unmarshal(data, &users)
	assert.NoError(t, err)
	assert.Len(t, users, 1)
	assert.Equal(t, newUser.Name, users[0].Name)
}
