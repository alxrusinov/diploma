package usecase

import "github.com/alxrusinov/diploma/internal/model"

func (useCase *Usecase) GetWithdrawls(login string) ([]model.Balance, error) {
	return make([]model.Balance, 0), nil
}
