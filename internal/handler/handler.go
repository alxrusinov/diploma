package handler

import (
	"github.com/alxrusinov/diploma/internal/model"
)

type Usecase interface {
	CheckUserExists(user *model.User) (bool, error)
	CreateUser(user *model.User) (string, error)
	UpdateUser(token *model.Token) (*model.Token, error)
	CheckIsValidUser(user *model.User) (string, error)
	UploadOrder(order *model.Order, userID string) error
	GetOrders(login string) ([]model.OrderResponse, error)
	GetBalance(userID string) (*model.Balance, error)
	GetOrder(order *model.Order, userID string) (*model.Order, error)
	GetWithdrawls(userID string) ([]model.Withdrawn, error)
	SetWithdrawls(withdrawn *model.Withdrawn, userID string) error
	UpdateBalance(balance float64, userID string) error
}

type Auth interface {
	GetToken(user *model.User) (*model.Token, error)
	ParseToken(tokenString string) (*model.Token, error)
}

type Handler struct {
	usecase    Usecase
	Middleware Middleware
	AuthClient Auth
}

const (
	TokenCookie = "token"
)

func NewHandler(usecaseInst Usecase, responseAddr string, authClient Auth) *Handler {
	handler := &Handler{
		usecase:    usecaseInst,
		Middleware: Middleware{},
		AuthClient: authClient,
	}

	return handler
}
