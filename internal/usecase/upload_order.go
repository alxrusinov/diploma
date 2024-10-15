package usecase

import (
	"context"
	"errors"

	"github.com/alxrusinov/diploma/internal/customerrors"
	"github.com/alxrusinov/diploma/internal/model"
)

func (usecase *Usecase) UploadOrder(order *model.Order, userID string) error {
	noOrderError := new(customerrors.NoOrderError)

	orderUserID, err := usecase.store.CheckOrder(order)

	if err != nil {
		if errors.As(err, &noOrderError) {
			order.Process = model.New
			_, err = usecase.store.AddOrder(order, userID)

			if err != nil {
				return err
			}

		} else {
			return err

		}

	}

	if err == nil && orderUserID != "" {
		if orderUserID == userID {
			return &customerrors.DuplicateOwnerOrderError{}
		}

		return &customerrors.DuplicateUserOrderError{}
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
