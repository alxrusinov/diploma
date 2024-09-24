package handler

import (
	"github.com/alxrusinov/diploma/internal/app/useCase"
	"github.com/alxrusinov/diploma/internal/auth"
	"github.com/alxrusinov/diploma/internal/client"
)

type options struct {
	responseAddr string
}

type Handler struct {
	useCase    useCase.UseCase
	options    options
	client     *client.Client
	Middleware Middleware
	AuthClient *auth.Auth
}

const (
	TokenCookie = "token"
)

func CreateHandler(useCase useCase.UseCase, responseAddr string, authClient *auth.Auth) *Handler {
	handler := &Handler{
		useCase: useCase,
		options: options{
			responseAddr: responseAddr,
		},
		client:     client.CreateClient(),
		Middleware: Middleware{},
		AuthClient: authClient,
	}

	return handler
}
