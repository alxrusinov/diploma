package handler

import (
	"github.com/alxrusinov/diploma/internal/model"
)

type options struct {
	responseAddr string
}

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

type Auth interface {
	GetToken(user *model.User) (*model.Token, error)
	ParseToken(tokenString string) (*model.Token, error)
}

type Handler struct {
	usecase    Usecase
	options    options
	Middleware Middleware
	AuthClient Auth
}

const (
	TokenCookie = "token"
)

func NewHandler(usecaseInst Usecase, responseAddr string, authClient Auth) *Handler {
	handler := &Handler{
		usecase: usecaseInst,
		options: options{
			responseAddr: responseAddr,
		},
		Middleware: Middleware{},
		AuthClient: authClient,
	}

	return handler
}
