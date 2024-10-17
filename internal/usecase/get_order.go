package usecase

import "github.com/alxrusinov/diploma/internal/model"

func (usecase *Usecase) GetOrder(order *model.Order, userID string) (*model.Order, error) {

	order, err := usecase.store.GetOrder(order, userID)

	return order, err
}
