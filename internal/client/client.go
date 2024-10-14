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

func (client *Client) GetOrderInfo(ctx context.Context, orderNumber string, resCh chan<- *model.Order, cancel context.CancelFunc) {

	addr := fmt.Sprintf("%s/api/orders/%s", client.addr, orderNumber)

	tick := time.NewTicker(time.Second * 3)

	for {
		select {
		case <-ctx.Done():
			return
		case <-tick.C:
			req, err := http.NewRequest(http.MethodGet, addr, nil)

			if err != nil {
				cancel()
			}

			res, err := client.client.Do(req)

			if err != nil {
				cancel()
			}

			defer res.Body.Close()

			if res.StatusCode == http.StatusOK {
				order := new(model.Order)

				if err := json.NewDecoder(res.Body).Decode(order); err != nil && !errors.Is(err, io.EOF) {
					cancel()
				}

				resCh <- order
				return

			}

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
