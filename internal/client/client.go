package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/alxrusinov/diploma/internal/customerrors"
	"github.com/alxrusinov/diploma/internal/model"
)

type Client struct {
	client *http.Client
}

func (client *Client) GetOrderInfo(orderNumber string) (*model.Order, error) {

	addr := fmt.Sprintf("/api/orders/%s", orderNumber)

	req, err := http.NewRequest(http.MethodGet, addr, nil)

	if err != nil {
		return nil, err
	}

	res, err := client.client.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		order := &model.Order{}

		if err := json.NewDecoder(res.Request.Body).Decode(&order); err != nil && !errors.Is(err, io.EOF) {
			return nil, err
		}

		return order, nil
	}

	if res.StatusCode == http.StatusNoContent {
		return nil, &customerrors.NoOrderError{}
	}

	return nil, &customerrors.ServerError{}

}

func CreateClient() *Client {
	return &Client{
		client: &http.Client{},
	}
}
