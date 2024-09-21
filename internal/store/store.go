package store

import (
	"github.com/alxrusinov/diploma/internal/config"
	"github.com/alxrusinov/diploma/internal/model"
)

type Store interface {
	CheckUserExists(user *model.User) (bool, error)
	CreateUser(user *model.User) error
	UpdateUser(token *model.Token) (*model.Token, error)
}

func CreateStore(config *config.Config) Store {
	return CreateDBStore(config.DatabaseURI)
}
