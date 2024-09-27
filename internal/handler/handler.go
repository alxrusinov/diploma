package handler

import (
	"github.com/alxrusinov/diploma/internal/authenticate"
	"github.com/alxrusinov/diploma/internal/useCase"
)

type options struct {
	responseAddr string
}

type Handler struct {
	useCase    useCase.UseCase
	options    options
	Middleware Middleware
	AuthClient *authenticate.Auth
}

const (
	TokenCookie = "token"
)

func CreateHandler(useCase useCase.UseCase, responseAddr string, authClient *authenticate.Auth) *Handler {
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
