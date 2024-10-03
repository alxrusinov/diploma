package usecase

import "github.com/alxrusinov/diploma/internal/model"

func (usecase *Usecase) GetBalance(userID string) (*model.Balance, error) {
	balance, err := usecase.store.GetBalance(userID)

	if err != nil {
		return nil, err
	}

	return balance, nil
}
