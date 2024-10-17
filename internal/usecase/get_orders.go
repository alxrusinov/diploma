package usecase

import (
	"github.com/alxrusinov/diploma/internal/model"
)

func (usecase *Usecase) GetOrders(login string) ([]model.OrderResponse, error) {
	orders, err := usecase.store.GetOrders(login)

	if err != nil {
		return nil, err
	}

	return orders, nil
}
