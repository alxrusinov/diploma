package client

import (
	"net/http"

	"github.com/alxrusinov/diploma/internal/model"
)

type Client struct {
	client *http.Client
}

func (client *Client) GetOrderInfo(order string) (*model.Order, error) {
	return nil, nil
}

func CreateClient() *Client {
	return &Client{
		client: &http.Client{},
	}
}
