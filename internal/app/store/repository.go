package store

import (
	"github.com/bruhlord-s/virttable-api/internal/app/model"
	"github.com/google/uuid"
)

type UserRepository interface {
	Create(*model.User)    error
	Find(uuid.UUID)		   (*model.User, error)
	FindByUsername(string) (*model.User, error)
	FindByEmail(string)	   (*model.User, error)
}