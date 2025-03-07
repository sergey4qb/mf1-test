package user

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/sergey4qb/mf1-test/dto"
	"github.com/sergey4qb/mf1-test/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockRepo struct {
	users     []model.User
	createErr error
}

func (r *mockRepo) Create(ctx context.Context, user *model.User) error {
	if r.createErr != nil {
		return r.createErr
	}
	r.users = append(r.users, *user)
	return nil
}

func (r *mockRepo) GetAll(ctx context.Context) ([]model.User, error) {
	return r.users, nil
}

func (r *mockRepo) GetByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	for i := range r.users {
		if r.users[i].ID == id {
			return &r.users[i], nil
		}
	}
	return nil, errors.New("user not found")
}

func (r *mockRepo) Update(ctx context.Context, user *model.User) error {
	for i, u := range r.users {
		if u.ID == user.ID {
			r.users[i] = *user
			return nil
		}
	}
	return errors.New("user not found")
}

func (r *mockRepo) Delete(ctx context.Context, id uuid.UUID) error {
	for i, u := range r.users {
		if u.ID == id {
			r.users = append(r.users[:i], r.users[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}

func TestCreate_InvalidName(t *testing.T) {
	repo := &mockRepo{}
	srv := New(repo)
	u := &model.User{
		Name:  "",
		Email: "test@example.com",
	}
	err := srv.Create(context.Background(), u)
	assert.Equal(t, errInvalidName, err, "empty name should return error")
}

func TestCreate_InvalidEmail_Empty(t *testing.T) {
	repo := &mockRepo{}
	srv := New(repo)
	u := &model.User{
		Name:  "Test",
		Email: "",
	}
	err := srv.Create(context.Background(), u)
	assert.Equal(t, errInvalidEmail, err, "empty email should return error")
}

func TestCreate_InvalidEmail_Format(t *testing.T) {
	repo := &mockRepo{}
	srv := New(repo)
	u := &model.User{
		Name:  "Test",
		Email: "invalid-email",
	}
	err := srv.Create(context.Background(), u)
	assert.Equal(t, errInvalidFormatEmail, err, "invalid email should return error")
}

func TestCreate_Success(t *testing.T) {
	repo := &mockRepo{}
	srv := New(repo)
	u := &model.User{
		Name:  "Test",
		Email: "test@example.com",
	}
	err := srv.Create(context.Background(), u)
	assert.NoError(t, err, "correct data should not error")
	assert.NotEqual(t, uuid.Nil, u.ID, "user id should not be empty")
	users, err := repo.GetAll(context.Background())
	assert.NoError(t, err)
	assert.Len(t, users, 1)
	assert.Equal(t, u.Name, users[0].Name)
	assert.Equal(t, u.Email, users[0].Email)
}

func TestGetAll(t *testing.T) {
	repo := &mockRepo{}
	user1 := model.User{
		ID:    uuid.New(),
		Name:  "User1",
		Email: "user1@example.com",
	}
	user2 := model.User{
		ID:    uuid.New(),
		Name:  "User2",
		Email: "user2@example.com",
	}
	repo.users = []model.User{user1, user2}

	srv := New(repo)
	users, err := srv.GetAll(context.Background())
	assert.NoError(t, err)
	assert.Len(t, users, 2)
	assert.Equal(t, user1.Name, users[0].Name)
	assert.Equal(t, user2.Email, users[1].Email)
}

func TestGetByID_Success(t *testing.T) {
	repo := &mockRepo{}
	u := model.User{
		ID:    uuid.New(),
		Name:  "Test User",
		Email: "test@example.com",
	}
	repo.users = append(repo.users, u)

	srv := New(repo)
	userFromService, err := srv.GetByID(context.Background(), u.ID)
	assert.NoError(t, err)
	assert.Equal(t, u.ID, userFromService.ID)
	assert.Equal(t, u.Name, userFromService.Name)
	assert.Equal(t, u.Email, userFromService.Email)
}

func TestUpdate_Success(t *testing.T) {
	repo := &mockRepo{}
	u := model.User{
		ID:    uuid.New(),
		Name:  "Old Name",
		Email: "old@example.com",
	}
	repo.users = append(repo.users, u)

	srv := New(repo)
	newName := "New Name"
	newEmail := "new@example.com"
	updateDTO := &dto.UpdateUserDTO{
		ID:    u.ID,
		Name:  &newName,
		Email: &newEmail,
	}

	updatedUser, err := srv.Update(context.Background(), updateDTO)
	assert.NoError(t, err)
	assert.Equal(t, newName, updatedUser.Name)
	assert.Equal(t, newEmail, updatedUser.Email)

	us, err := srv.GetByID(context.Background(), updateDTO.ID)
	assert.NoError(t, err)
	assert.Equal(t, us.Name, updatedUser.Name)
	assert.Equal(t, us.Email, updatedUser.Email)
}

func TestDelete_Success(t *testing.T) {
	repo := &mockRepo{}
	u := model.User{
		ID:    uuid.New(),
		Name:  "Test User",
		Email: "test@example.com",
	}
	repo.users = append(repo.users, u)

	srv := New(repo)
	err := srv.Delete(context.Background(), u.ID)
	assert.NoError(t, err)

	_, err = srv.GetByID(context.Background(), u.ID)
	assert.Error(t, err)
}
