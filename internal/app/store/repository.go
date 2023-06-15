package store

import "github.com/bruhlord-s/virttable-api/internal/app/model"

type UserRepository interface {
	Create(*model.User)    error
	FindByUsername(string) (*model.User, error)
}