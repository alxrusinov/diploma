package usecase

import "github.com/alxrusinov/diploma/internal/model"

func (usecase *Usecase) SetWithdrawls(withdrawn *model.Withdrawn, userID string) error {
	err := usecase.store.SetWithdrawls(withdrawn, userID)

	return err
}
