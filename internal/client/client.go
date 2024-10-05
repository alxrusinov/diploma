package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

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
			order := new(model.Order)

			if err := json.NewDecoder(res.Body).Decode(order); err != nil && !errors.Is(err, io.EOF) {
				return nil, err
			}

			if order.Process == "PROCESSED" {
				order.Round()
				return order, nil
			}

			if order.Process == "INVALID" {
				return nil, errors.New("invalid order")
			}

		}

		if res.StatusCode == http.StatusTooManyRequests {
			return nil, errors.New("too many requests")
		}

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
