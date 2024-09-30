package usecase

import "github.com/alxrusinov/diploma/internal/model"

func (useCase *Usecase) GetBalance(login string) (*model.Balance, error) {
	order, err := useCase.client.GetOrderInfo(login)

	if err != nil {
		return nil, err
	}

	balance := &model.Balance{
		Current: float64(order.Accrual),
	}

	return balance, nil
}
