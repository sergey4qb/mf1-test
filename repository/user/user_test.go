package user

import (
	"context"
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

	users, err := repo.GetAll(context.Background())
	assert.NoError(t, err)
	assert.Len(t, users, 0)

	newUser := &model.User{
		ID:    uuid.New(),
		Name:  "Test User",
		Email: "test@example.com",
	}
	err = repo.Create(context.Background(), newUser)
	assert.NoError(t, err)

	users, err = repo.GetAll(context.Background())
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

	users, err := repo.GetAll(context.Background())
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

	users, err := repo.GetAll(context.Background())
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
	err = repo.Create(context.Background(), newUser)
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
	err = repo.Create(context.Background(), newUser)
	assert.NoError(t, err)

	data, err := os.ReadFile(fileRepoPath)
	assert.NoError(t, err)

	var users []model.User
	err = json.Unmarshal(data, &users)
	assert.NoError(t, err)
	assert.Len(t, users, 1)
	assert.Equal(t, newUser.Name, users[0].Name)
}

func TestFileUserRepository_GetByID_Success(t *testing.T) {
	tempDir := t.TempDir()
	fileRepoPath = filepath.Join(tempDir, "users.json")
	_ = os.Remove(fileRepoPath)

	repo, err := New()
	assert.NoError(t, err)

	newUser := &model.User{
		ID:    uuid.New(),
		Name:  "GetByID User",
		Email: "getbyid@example.com",
	}
	err = repo.Create(context.Background(), newUser)
	assert.NoError(t, err)

	foundUser, err := repo.GetByID(context.Background(), newUser.ID)
	assert.NoError(t, err)
	assert.Equal(t, newUser.ID, foundUser.ID)
	assert.Equal(t, newUser.Name, foundUser.Name)
	assert.Equal(t, newUser.Email, foundUser.Email)
}

func TestFileUserRepository_GetByID_NotFound(t *testing.T) {
	tempDir := t.TempDir()
	fileRepoPath = filepath.Join(tempDir, "users.json")
	_ = os.Remove(fileRepoPath)

	repo, err := New()
	assert.NoError(t, err)

	nonExistentID := uuid.New()
	user, err := repo.GetByID(context.Background(), nonExistentID)
	assert.Error(t, err)
	assert.Nil(t, user)
}

func TestFileUserRepository_Update_Success(t *testing.T) {
	tempDir := t.TempDir()
	fileRepoPath = filepath.Join(tempDir, "users.json")
	_ = os.Remove(fileRepoPath)

	repo, err := New()
	assert.NoError(t, err)

	origUser := &model.User{
		ID:    uuid.New(),
		Name:  "Original Name",
		Email: "original@example.com",
	}
	err = repo.Create(context.Background(), origUser)
	assert.NoError(t, err)

	origUser.Name = "Updated Name"
	origUser.Email = "updated@example.com"
	err = repo.Update(context.Background(), origUser)
	assert.NoError(t, err)

	updatedUser, err := repo.GetByID(context.Background(), origUser.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Updated Name", updatedUser.Name)
	assert.Equal(t, "updated@example.com", updatedUser.Email)
}

func TestFileUserRepository_Update_NotFound(t *testing.T) {
	tempDir := t.TempDir()
	fileRepoPath = filepath.Join(tempDir, "users.json")
	_ = os.Remove(fileRepoPath)

	repo, err := New()
	assert.NoError(t, err)

	nonExistentUser := &model.User{
		ID:    uuid.New(),
		Name:  "NonExistent",
		Email: "nonexistent@example.com",
	}
	err = repo.Update(context.Background(), nonExistentUser)
	assert.Error(t, err)
}

func TestFileUserRepository_Delete_Success(t *testing.T) {
	tempDir := t.TempDir()
	fileRepoPath = filepath.Join(tempDir, "users.json")
	_ = os.Remove(fileRepoPath)

	repo, err := New()
	assert.NoError(t, err)

	userToDelete := &model.User{
		ID:    uuid.New(),
		Name:  "Delete Me",
		Email: "delete@example.com",
	}
	err = repo.Create(context.Background(), userToDelete)
	assert.NoError(t, err)

	err = repo.Delete(context.Background(), userToDelete.ID)
	assert.NoError(t, err)

	_, err = repo.GetByID(context.Background(), userToDelete.ID)
	assert.Error(t, err)
}

func TestFileUserRepository_Delete_NotFound(t *testing.T) {
	tempDir := t.TempDir()
	fileRepoPath = filepath.Join(tempDir, "users.json")
	_ = os.Remove(fileRepoPath)

	repo, err := New()
	assert.NoError(t, err)

	err = repo.Delete(context.Background(), uuid.New())
	assert.Error(t, err)
}
