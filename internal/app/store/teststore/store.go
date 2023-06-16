package teststore

import (
	"github.com/bruhlord-s/virttable-api/internal/app/model"
	"github.com/bruhlord-s/virttable-api/internal/app/store"
)

type Store struct {
	UserRepository *UserRepository
}

func New() *Store {
	return &Store{}
}

func (s *Store) User() store.UserRepository {
	if s.UserRepository != nil {
		return s.UserRepository
	}

	s.UserRepository = &UserRepository{
		store: s,
		usersByUsername: make(map[string]*model.User),
		usersByEmail: make(map[string]*model.User),
	}

	return s.UserRepository
}