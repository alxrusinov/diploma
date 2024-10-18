package usecase

import (
	"errors"

	"github.com/alxrusinov/diploma/internal/customerrors"
	"github.com/alxrusinov/diploma/internal/logger"
	"github.com/alxrusinov/diploma/internal/model"
	"go.uber.org/zap"
)

func (usecase *Usecase) UploadOrder(order *model.Order, userID string) error {
	noOrderError := new(customerrors.NoOrderError)

	orderUserID, err := usecase.store.CheckOrder(order)

	if err != nil {
		if errors.As(err, &noOrderError) {
			order.Process = model.New
			_, err = usecase.store.AddOrder(order, userID)
			logger.Logger.Info("order added", zap.String("order number", order.Number))

			if err != nil {
				return err
			}

		} else {
			return err

		}

	}

	if err == nil && orderUserID != "" {
		if orderUserID == userID {
			return &customerrors.DuplicateOwnerOrderError{Err: errors.New("current user has already added order")}
		}

		return &customerrors.DuplicateUserOrderError{Err: errors.New("another user has already added order")}
	}

	return nil
}
