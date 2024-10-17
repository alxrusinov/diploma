package usecase

import (
	"errors"

	"github.com/alxrusinov/diploma/internal/customerrors"
	"github.com/alxrusinov/diploma/internal/model"
)

func (usecase *Usecase) GetBalance(userID string) (*model.Balance, error) {
	balance, err := usecase.store.GetBalance(userID)

	if err != nil {
		return nil, err
	}

	withdrawls, err := usecase.GetWithdrawls(userID)

	if err != nil {
		noContentError := new(customerrors.NoContentError)

		if errors.As(err, &noContentError) {
			balance.Withdrawn = 0
			return balance, nil
		}

		return nil, err
	}

	for _, withdrawn := range withdrawls {
		balance.Withdrawn = balance.Withdrawn + withdrawn.Sum
	}

	return balance, nil
}
