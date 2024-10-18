package orderclient

import (
	"context"
	"time"

	"github.com/alxrusinov/diploma/internal/logger"
	"github.com/alxrusinov/diploma/internal/model"
	"go.uber.org/zap"
)

type Store interface {
	GetProcessingOrder(ctx context.Context) ([]model.Order, error)
	UpdateOrder(ctx context.Context, order *model.Order) error
}

type OrderClient struct {
	store Store
}

func (orderClient *OrderClient) GetProcessingOrder(ctx context.Context) <-chan *model.Order {
	orderCh := make(chan *model.Order)

	ticker := time.NewTicker(time.Second)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				res, err := orderClient.store.GetProcessingOrder(ctx)

				if err != nil {
					logger.Logger.Error("get orders error", zap.Error(err))
					continue
				}

				logger.Logger.Info("get orders", zap.Any("orders", res))

				for _, order := range res {
					logger.Logger.Info("order to chan", zap.Any("order", order))
					orderCh <- &order
				}
			}
		}
	}()

	return orderCh

}

func (orderClient *OrderClient) UpdateOrder(ctx context.Context, inChan <-chan *model.Order) {

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case order := <-inChan:
				err := orderClient.store.UpdateOrder(ctx, order)

				if err != nil {
					logger.Logger.Error("error upload order", zap.Error(err), zap.Any("ORDER", order))
					continue
				}
			}
		}
	}()

}

func NewOrderClient(store Store) *OrderClient {
	return &OrderClient{
		store: store,
	}
}
