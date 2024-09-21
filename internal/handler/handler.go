package handler

import (
	auth "github.com/alxrusinov/diploma/internal/Auth"
	"github.com/alxrusinov/diploma/internal/client"
	"github.com/alxrusinov/diploma/internal/store"
)

type options struct {
	responseAddr string
}

type Handler struct {
	store      store.Store
	options    options
	client     *client.Client
	Middleware Middleware
	AuthClient *auth.Auth
}

const (
	TokenCookie = "token"
)

func CreateHandler(currentStore store.Store, responseAddr string, authClient *auth.Auth) *Handler {
	handler := &Handler{
		store: currentStore,
		options: options{
			responseAddr: responseAddr,
		},
		client:     client.CreateClient(),
		Middleware: Middleware{},
		AuthClient: authClient,
	}

	return handler
}
