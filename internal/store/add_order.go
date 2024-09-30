package store

import (
	"time"

	"github.com/alxrusinov/diploma/internal/model"
)

func (store *Store) AddOrder(order *model.Order, login string) (bool, error) {
	tx, err := store.db.Begin()

	if err != nil {
		return false, err
	}

	var userID string

	err = tx.QueryRow(selectUserByLoginQuery, login).Scan(&userID)

	if err != nil {
		tx.Rollback()
		return false, err
	}

	row, err := tx.Query(insertOrderQuery, userID, order.Number, order.Process, order.Accrual, time.Now().Format(time.RFC3339))

	if err != nil || row.Err() != nil {
		tx.Rollback()
		return false, err
	}

	var sum int

	err = tx.QueryRow(selectBalanceQuery, userID).Scan(&sum)

	if err != nil {
		tx.Rollback()
		return false, err
	}

	rows, err := tx.Query(updateBalanceQuery, sum+order.Accrual, userID)

	if err != nil || rows.Err() != nil {
		tx.Rollback()
		return false, err
	}

	tx.Commit()

	return true, nil
}
