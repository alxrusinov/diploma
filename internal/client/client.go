package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/alxrusinov/diploma/internal/customerrors"
	"github.com/alxrusinov/diploma/internal/model"
)

type Client struct {
	client *http.Client
	addr   string
}

func (client *Client) GetOrderInfo(orderNumber string) (*model.Order, error) {

	addr := fmt.Sprintf("%s/api/orders/%s", client.addr, orderNumber)

	for {
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

		return nil, &customerrors.ServerError{}
	}

}

func NewClient(addr string) *Client {
	return &Client{
		client: &http.Client{
			Timeout: time.Second * 60,
		},
		addr: addr,
	}
}
