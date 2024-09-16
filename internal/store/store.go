package store

import "github.com/alxrusinov/diploma/internal/config"

type Store interface {
}

func CreateStore(config *config.Config) Store {
	return CreateDBStore(config.DatabaseURI)
}
