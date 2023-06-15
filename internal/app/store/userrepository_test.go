package store_test

import (
	"testing"

	"github.com/bruhlord-s/virttable-api/internal/app/model"
	"github.com/bruhlord-s/virttable-api/internal/app/store"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	u, err := s.User().Create(&model.User{
		Email: "test@test.com",
		Username: "test",
	})
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByUsername(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	username := "test"
	_, err := s.User().FindByUsername(username)
	assert.Error(t, err)

	s.User().Create(&model.User{
		Username: "test",
	})
	u, err := s.User().FindByUsername(username)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}