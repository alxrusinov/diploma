package usecase

import (
	"context"
	"errors"

	"github.com/alxrusinov/diploma/internal/customerrors"
	"github.com/alxrusinov/diploma/internal/logger"
	"github.com/alxrusinov/diploma/internal/model"
	"go.uber.org/zap"
)

func (usecase *Usecase) UploadOrder(order *model.Order, userID string) error {
	noOrderError := new(customerrors.NoOrderError)

	logger.Logger.Info("start order uploading", zap.String("order number", order.Number))
	orderUserID, err := usecase.store.CheckOrder(order)

	if err != nil {
		if errors.As(err, &noOrderError) {
			order.Process = model.New
			_, err = usecase.store.AddOrder(order, userID)
			logger.Logger.Info("order added", zap.String("order number", order.Number))

			if err != nil {
				logger.Logger.Error("error add order", zap.Error(err), zap.String("order number", order.Number))
				return err
			}

		} else {
			logger.Logger.Error("error check order", zap.Error(err), zap.String("order number", order.Number))
			return err

		}

	}

	if err == nil && orderUserID != "" {
		if orderUserID == userID {
			return &customerrors.DuplicateOwnerOrderError{Err: errors.New("current user has already added order")}
		}

		return &customerrors.DuplicateUserOrderError{Err: errors.New("another user has already added order")}
	}

	ctx, cancel := context.WithCancel(context.Background())

	orderCh, clienErr := usecase.client.GetOrderInfo(ctx, order.Number)

	done, errCh := usecase.store.UpdateOrder(ctx, userID, orderCh)

	go func() {
		select {
		case <-done:
			cancel()
			return
		case <-clienErr:
			cancel()
			return
		case <-errCh:
			cancel()
			return
		}
	}()

	return nil
}
