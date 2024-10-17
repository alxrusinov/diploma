package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/alxrusinov/diploma/internal/logger"
	"github.com/alxrusinov/diploma/internal/model"
	"go.uber.org/zap"
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
					logger.Logger.Error("client: create request error", zap.Error(err))
					close(errCh)
				}

				res, err := client.client.Do(req)

				if err != nil {

					if res.StatusCode == http.StatusTooManyRequests {
						logger.Logger.Error("client: too many requests", zap.Error(err))
						time.Sleep(time.Second * 5)
						continue
					}

					logger.Logger.Error("client: another error", zap.Error(err), zap.Any("response", res))
					close(errCh)
				}

				defer res.Body.Close()

				if res.StatusCode == http.StatusOK {
					order := new(model.Order)

					if err := json.NewDecoder(res.Body).Decode(order); err != nil && !errors.Is(err, io.EOF) {
						logger.Logger.Error("client: unmarshaling error", zap.Error(err), zap.Any("response", res))
						close(errCh)
						return
					}

					logger.Logger.Info("client: success get order", zap.Any("response", res), zap.Any("order", order))
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
