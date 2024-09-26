package useCase

import (
	"github.com/alxrusinov/diploma/internal/model"
)

type UseCase interface {
	CheckUserExists(user *model.User) (bool, error)
	CreateUser(user *model.User) error
	UpdateUser(token *model.Token) (*model.Token, error)
	CheckIsValidUser(user *model.User) (bool, error)
	UploadOrder(order *model.Order) (*model.Order, error)
	GetOrders(login string) ([]model.OrderResponse, error)
	GetBalance(login string) (*model.Balance, error)
}
