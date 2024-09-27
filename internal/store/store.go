package store

import (
	"github.com/alxrusinov/diploma/internal/config"
	"github.com/alxrusinov/diploma/internal/model"
)

type Store interface {
	FindUserByLogin(user *model.User) (bool, error)
	FindUserByLoginPassword(user *model.User) (bool, error)
	CreateUser(user *model.User) (bool, error)
	UpdateUser(token *model.Token) (*model.Token, error)
	AddOrder(order *model.Order) (bool, error)
	GetOrders(login string) ([]model.OrderResponse, error)
}

func CreateStore(config *config.Config) Store {
	return CreateDBStore(config.DatabaseURI)
}
