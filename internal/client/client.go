package client

import (
	"context"
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

func (client *Client) GetOrderInfo(ctx context.Context, orderNumber string) (<-chan *model.Order, <-chan error) {

	addr := fmt.Sprintf("%s/api/orders/%s", client.addr, orderNumber)

	orderCh := make(chan *model.Order)
	errCh := make(chan error)

	go func() {
		tick := time.NewTicker(time.Second)

		defer tick.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-tick.C:
				req, err := http.NewRequest(http.MethodGet, addr, nil)

				if err != nil {
					close(errCh)
				}

				res, err := client.client.Do(req)

				if err != nil {

					if res.StatusCode == http.StatusTooManyRequests {
						time.Sleep(time.Second * 5)
						continue
					}

					close(errCh)
				}

				defer res.Body.Close()

				if res.StatusCode == http.StatusOK {
					order := new(model.Order)

					if err := json.NewDecoder(res.Body).Decode(order); err != nil && !errors.Is(err, io.EOF) {
						close(errCh)
						return
					}

					orderCh <- order
					return

				}

			}

		}
	}()

	return orderCh, errCh

}

func NewClient(addr string, timeout time.Duration) *Client {
	return &Client{
		client: &http.Client{
			Timeout: timeout,
		},
		addr: addr,
	}
}
