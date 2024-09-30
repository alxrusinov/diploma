package usecase

import (
	"errors"

	"github.com/alxrusinov/diploma/internal/customerrors"
	"github.com/alxrusinov/diploma/internal/model"
)

func (useCase *Usecase) UploadOrder(order *model.Order, login string) (*model.Order, error) {
	var resOrder *model.Order
	noOrderError := new(customerrors.NoOrderError)
	serverError := new(customerrors.ServerError)

	for {
		resOrder, err := useCase.client.GetOrderInfo(order.Number)

		if err != nil {
			if errors.As(err, &serverError) {
				continue
			}

			if errors.As(err, &noOrderError) {
				return nil, err
			}

			return nil, err
		}

		ok, err := useCase.store.AddOrder(resOrder, login)

		if err != nil {
			return nil, err
		}

		if !ok {
			return nil, errors.New("order was not uploaded")
		}

		break

	}

	return resOrder, nil
}
