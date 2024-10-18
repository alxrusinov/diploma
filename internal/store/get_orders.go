package store

import (
	"github.com/alxrusinov/diploma/internal/logger"
	"github.com/alxrusinov/diploma/internal/model"
	"go.uber.org/zap"
)

func (store *Store) GetOrders(login string) ([]model.OrderResponse, error) {
	res := make([]model.OrderResponse, 0)
	var userID string

	err := store.db.QueryRow(selectUserByLoginQuery, login).Scan(&userID)

	if err != nil {
		logger.Logger.Error("error user selection", zap.Error(err))
		return res, err
	}

	rows, err := store.db.Query(selectOrdersQuery, userID)

	if err != nil {
		logger.Logger.Error("error selection orders", zap.Error(err))
		return res, err
	}

	defer rows.Close()

	for rows.Next() {
		item := new(model.OrderResponse)
		err = rows.Scan(&item.Number, &item.Status, &item.Accrual, &item.UploadedAt)

		if err != nil {
			logger.Logger.Error("scan error", zap.Error(err))
			return res, err
		}

		item.Round()

		res = append(res, *item)

	}

	if err := rows.Err(); err != nil {
		logger.Logger.Error("rows error", zap.Error(err))
		return res, err
	}

	logger.Logger.Info("success orders", zap.Any("response", res))

	return res, nil
}
