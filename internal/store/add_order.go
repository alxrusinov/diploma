package store

import (
	"time"

	"github.com/alxrusinov/diploma/internal/logger"
	"github.com/alxrusinov/diploma/internal/mathfn"
	"github.com/alxrusinov/diploma/internal/model"
	"go.uber.org/zap"
)

func (store *Store) AddOrder(order *model.Order, userID string) (bool, error) {
	tx, err := store.db.Begin()

	if err != nil {
		return false, err
	}

	order.Round()

	row, err := tx.Query(insertOrderQuery, userID, order.Number, order.Process, order.Accrual, time.Now().Format(time.RFC3339))

	logger.Logger.Info("time format", zap.String("time", time.Now().Format(time.RFC3339)))

	if err != nil || row.Err() != nil {
		tx.Rollback()
		return false, err
	}

	var sum float64

	err = tx.QueryRow(selectBalanceQuery, userID).Scan(&sum)

	if err != nil {
		tx.Rollback()
		return false, err
	}

	rows, err := tx.Query(updateBalanceQuery, mathfn.RoundFloat(sum+order.Accrual, 5), userID)

	if err != nil || rows.Err() != nil {
		tx.Rollback()
		return false, err
	}

	tx.Commit()

	return true, nil
}
