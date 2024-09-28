package handler

import (
	"github.com/alxrusinov/diploma/internal/authenticate"
	"github.com/alxrusinov/diploma/internal/usecase"
)

type options struct {
	responseAddr string
}

type Handler struct {
	useCase    usecase.Usecase
	options    options
	Middleware Middleware
	AuthClient *authenticate.Auth
}

const (
	TokenCookie = "token"
)

func CreateHandler(useCase usecase.Usecase, responseAddr string, authClient *authenticate.Auth) *Handler {
	handler := &Handler{
		useCase: useCase,
		options: options{
			responseAddr: responseAddr,
		},
		Middleware: Middleware{},
		AuthClient: authClient,
	}

	return handler
}
