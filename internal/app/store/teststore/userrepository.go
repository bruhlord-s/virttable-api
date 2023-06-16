package teststore

import (
	"github.com/bruhlord-s/virttable-api/internal/app/model"
	"github.com/bruhlord-s/virttable-api/internal/app/store"
	"github.com/google/uuid"
)

type UserRepository struct {
	store 		 	*Store
	usersByEmail 	map[string]*model.User
	usersByUsername map[string]*model.User
}

func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	r.usersByUsername[u.Username] = u
	r.usersByEmail[u.Email] = u
	u.ID = uuid.New()

	return nil
}

func (r *UserRepository) FindByUsername(username string) (*model.User, error) {
	u, ok := r.usersByUsername[username]
	if !ok {
		return nil, store.ErrRecordNotFound
	}

	return u, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u, ok := r.usersByEmail[email]
	if !ok {
		return nil, store.ErrRecordNotFound
	}

	return u, nil
}