package teststore_test

import (
	"testing"

	"github.com/bruhlord-s/virttable-api/internal/app/model"
	"github.com/bruhlord-s/virttable-api/internal/app/store"
	"github.com/bruhlord-s/virttable-api/internal/app/store/teststore"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s := teststore.New()
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByUsername(t *testing.T) {
	s := teststore.New()

	username := "test"
	_, err := s.User().FindByUsername(username)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	tu := model.TestUser(t)
	s.User().Create(tu)
	u, err := s.User().FindByUsername(tu.Username)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s := teststore.New()

	email := "test@test.com"
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	tu := model.TestUser(t)
	s.User().Create(tu)
	u, err := s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_Find(t *testing.T) {
	s := teststore.New()

	id := uuid.New()
	_, err := s.User().Find(id)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	tu := model.TestUser(t)
	s.User().Create(tu)
	u, err := s.User().Find(tu.ID)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}