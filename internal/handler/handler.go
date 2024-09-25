package handler

import (
	"github.com/alxrusinov/diploma/internal/authenticate"
	"github.com/alxrusinov/diploma/internal/client"
	"github.com/alxrusinov/diploma/internal/useCase"
)

type options struct {
	responseAddr string
}

type Handler struct {
	useCase    useCase.UseCase
	options    options
	client     *client.Client
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
		client:     client.CreateClient(),
		Middleware: Middleware{},
		AuthClient: authClient,
	}

	return handler
}
