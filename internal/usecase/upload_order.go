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

	orderCh := make(chan *model.Order)

	go func(ctx context.Context, orderNumber string, resCh chan<- *model.Order, cancel context.CancelFunc) {
		usecase.client.GetOrderInfo(ctx, orderNumber, orderCh, cancel)

	}(ctx, order.Number, orderCh, cancel)

	go func(ctx context.Context, userID string, orderCh <-chan *model.Order, cancel context.CancelFunc) {
		usecase.store.UpdateOrder(ctx, userID, orderCh, cancel)
	}(ctx, userID, orderCh, cancel)

	return nil
}
