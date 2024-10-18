package usecase

import (
	"github.com/alxrusinov/diploma/internal/client"
	"github.com/alxrusinov/diploma/internal/model"
)

type Usecase struct {
	store  Store
	client *client.Client
}

type Store interface {
	FindUserByLogin(user *model.User) (bool, error)
	FindUserByLoginPassword(user *model.User) (string, error)
	CreateUser(user *model.User) (string, error)
	UpdateUser(token *model.Token) (*model.Token, error)
	AddOrder(order *model.Order, userID string) (bool, error)
	GetOrders(login string) ([]model.OrderResponse, error)
	RunMigration() error
	GetOrder(order *model.Order, userID string) (*model.Order, error)
	CheckOrder(order *model.Order) (string, error)
	GetBalance(userID string) (*model.Balance, error)
	UpdateBalance(balance float64, userID string) error
	SetWithdrawls(withdrawn *model.Withdrawn, userID string) error
	GetWithdrawls(userID string) ([]model.Withdrawn, error)
}

func NewUsecase(store Store, client *client.Client) *Usecase {
	return &Usecase{store: store, client: client}
}
