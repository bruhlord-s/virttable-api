package sqlstore_test

import (
	"testing"

	"github.com/bruhlord-s/virttable-api/internal/app/model"
	"github.com/bruhlord-s/virttable-api/internal/app/store"
	"github.com/bruhlord-s/virttable-api/internal/app/store/sqlstore"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByUsername(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)

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
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)

	email := "test@test.com"
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	tu := model.TestUser(t)
	s.User().Create(tu)
	u, err := s.User().FindByEmail(tu.Email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_Find(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)

	id := uuid.New()
	_, err := s.User().Find(id)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	tu := model.TestUser(t)
	s.User().Create(tu)
	u, err := s.User().Find(tu.ID)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}