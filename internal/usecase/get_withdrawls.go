package usecase

import "github.com/alxrusinov/diploma/internal/model"

func (usecase *Usecase) GetWithdrawls(userID string) ([]model.Withdrawn, error) {
	res, err := usecase.store.GetWithdrawls(userID)

	return res, err
}
