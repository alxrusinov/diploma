package usecase

import (
	"errors"

	"github.com/alxrusinov/diploma/internal/customerrors"
	"github.com/alxrusinov/diploma/internal/model"
)

func (useCase *Usecase) UploadOrder(order *model.Order, userID string) error {
	noOrderError := new(customerrors.NoOrderError)

	orderUserID, err := useCase.store.CheckOrder(order)

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
		resOrder, _ := useCase.client.GetOrderInfo(order.Number)

		useCase.store.AddOrder(resOrder, userID)
	}()

	return nil
}
