package handler

import (
	"github.com/alxrusinov/diploma/internal/client"
	"github.com/alxrusinov/diploma/internal/store"
)

type options struct {
	responseAddr string
}

type Handler struct {
	store      store.Store
	options    options
	Middleware Middleware
	client     *client.Client
}

func CreateHandler(currentStore store.Store, responseAddr string) *Handler {
	handler := &Handler{
		store: currentStore,
		options: options{
			responseAddr: responseAddr,
		},
		Middleware: Middleware{},
		client:     client.CreateClient(),
	}

	return handler
}
