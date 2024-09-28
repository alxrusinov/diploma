package handler

import (
	"github.com/alxrusinov/diploma/internal/authenticate"
	"github.com/alxrusinov/diploma/internal/use"
)

type options struct {
	responseAddr string
}

type Handler struct {
	usecase    use.Usecase
	options    options
	Middleware Middleware
	AuthClient *authenticate.Auth
}

const (
	TokenCookie = "token"
)

func CreateHandler(usecaseInst use.Usecase, responseAddr string, authClient *authenticate.Auth) *Handler {
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
