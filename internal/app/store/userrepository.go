package store

import "github.com/bruhlord-s/virttable-api/internal/app/model"

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	if err := u.Validate(); err != nil {
		return nil, err
	}

	if err := u.BeforeCreate(); err != nil {
		return nil, err
	}

	if err := r.store.db.QueryRow(
		"INSERT INTO users (email, username, encrypted_password) VALUES ($1, $2, $3) RETURNING id",
		u.Email,
		u.Username,
		u.EncryptedPassword,
	).Scan(&u.ID); err != nil {
		return nil, err
	}

	return u, nil
}

func (r *UserRepository) FindByUsername(username string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, email, username, encrypted_password FROM users WHERE username=$1",
		username,
	).Scan(&u.ID, &u.Email, &u.Username, &u.EncryptedPassword); err != nil {
		return nil, err
	}

	return u, nil
}