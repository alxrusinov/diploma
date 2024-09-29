package use

import (
	"github.com/alxrusinov/diploma/internal/model"
)

type Usecase interface {
	CheckUserExists(user *model.User) (bool, error)
	CreateUser(user *model.User) error
	UpdateUser(token *model.Token) (*model.Token, error)
	CheckIsValidUser(user *model.User) (bool, error)
	UploadOrder(order *model.Order, login string) (*model.Order, error)
	GetOrders(login string) ([]model.OrderResponse, error)
	GetBalance(login string) (*model.Balance, error)
	GetWithdrawls(login string) ([]model.Balance, error)
}
