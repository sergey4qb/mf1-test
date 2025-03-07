package user

import (
	"github.com/google/uuid"
	"github.com/sergey4qb/mf1-test/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockRepo struct {
	users     []model.User
	createErr error
}

func (r *mockRepo) Create(user *model.User) error {
	if r.createErr != nil {
		return r.createErr
	}
	r.users = append(r.users, *user)
	return nil
}

func (r *mockRepo) GetAll() ([]model.User, error) {
	return r.users, nil
}

func TestCreate_InvalidName(t *testing.T) {
	repo := &mockRepo{}
	srv := New(repo)
	u := &model.User{
		Name:  "",
		Email: "test@example.com",
	}
	err := srv.Create(u)
	assert.Equal(t, errInvalidName, err, "empty name should return error")
}

func TestCreate_InvalidEmail_Empty(t *testing.T) {
	repo := &mockRepo{}
	srv := New(repo)
	u := &model.User{
		Name:  "Test",
		Email: "",
	}
	err := srv.Create(u)
	assert.Equal(t, errInvalidEmail, err, "empty email should return error")
}

func TestCreate_InvalidEmail_Format(t *testing.T) {
	repo := &mockRepo{}
	srv := New(repo)
	u := &model.User{
		Name:  "Test",
		Email: "invalid-email",
	}
	err := srv.Create(u)
	assert.Equal(t, errInvalidFormatEmail, err, "invalid email should return error")
}

func TestCreate_Success(t *testing.T) {
	repo := &mockRepo{}
	srv := New(repo)
	u := &model.User{
		Name:  "Test",
		Email: "test@example.com",
	}
	err := srv.Create(u)
	assert.NoError(t, err, "correct data should not error")
	assert.NotEqual(t, uuid.Nil, u.ID, "user id should not be empty")
	users, err := repo.GetAll()
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
	users, err := srv.GetAll()
	assert.NoError(t, err)
	assert.Len(t, users, 2)
	assert.Equal(t, user1.Name, users[0].Name)
	assert.Equal(t, user2.Email, users[1].Email)
}
