package usecase

import (
	"errors"

	"github.com/alxrusinov/diploma/internal/customerrors"
	"github.com/alxrusinov/diploma/internal/model"
)

func (usecase *Usecase) UploadOrder(order *model.Order, userID string) error {
	noOrderError := new(customerrors.NoOrderError)

	orderUserID, err := usecase.store.CheckOrder(order)

	if err != nil && !errors.As(err, &noOrderError) {
		return err
	}

	if orderUserID != "" {
		if orderUserID == userID {
			return &customerrors.DuplicateOwnerOrderError{}
		}

		return &customerrors.DuplicateUserOrderError{}
	}

	go func() {
		if resOrder, err := usecase.client.GetOrderInfo(order.Number); err != nil {
			usecase.store.AddOrder(resOrder, userID)
		}

	}()

	return nil
}
