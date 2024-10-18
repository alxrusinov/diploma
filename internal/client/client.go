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

func (client *Client) GetOrderInfo(ctx context.Context, inChan <-chan *model.Order) <-chan *model.Order {

	orderCh := make(chan *model.Order)

	go func() {
		select {
		case <-ctx.Done():
			return
		default:
			for order := range inChan {
				addr := fmt.Sprintf("%s/api/orders/%s", client.addr, order.Number)

				logger.Logger.Info("get order", zap.Any("order", order))

				req, err := http.NewRequest(http.MethodGet, addr, nil)

				if err != nil {
					logger.Logger.Error("client: create request error", zap.Error(err))
					continue
				}

				res, err := client.client.Do(req)

				if err != nil {

					if res.StatusCode == http.StatusTooManyRequests {
						logger.Logger.Error("client: too many requests", zap.Error(err))
						time.Sleep(time.Second * 5)
						continue
					}

					logger.Logger.Error("client: another error", zap.Error(err), zap.Any("response", res))
					continue
				}

				if res.StatusCode == http.StatusOK {
					resOrder := new(model.Order)

					if err := json.NewDecoder(res.Body).Decode(resOrder); err != nil && !errors.Is(err, io.EOF) {
						logger.Logger.Error("client: unmarshaling error", zap.Error(err), zap.Any("response", res))
						res.Body.Close()
						continue
					}
					res.Body.Close()

					logger.Logger.Info("client: success get order", zap.Any("response", res), zap.Any("RESORDER", resOrder))

					resOrder.UserID = order.UserID
					orderCh <- resOrder
					continue

				}

			}
		}

	}()

	return orderCh

}

func NewClient(addr string, timeout time.Duration) *Client {
	return &Client{
		client: &http.Client{
			Timeout: timeout,
		},
		addr: addr,
	}
}
