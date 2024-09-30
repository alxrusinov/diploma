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

	err = tx.QueryRow(`SELECT id FROM users WHERE login = $1`, login).Scan(&userID)

	if err != nil {
		tx.Rollback()
		return false, err
	}

	row, err := tx.Query(`INSERT INTO orders (user_id, number, process, accrual, uploaded_at)
	VALUES ($1, $2, $3, $4, $5)`, userID, order.Number, order.Process, order.Accrual, time.Now().Format(time.RFC3339))

	if err != nil || row.Err() != nil {
		tx.Rollback()
		return false, err
	}

	var sum int

	err = tx.QueryRow(`SELECT balance FROM balance WHERE user_id = $1`, userID).Scan(&sum)

	if err != nil {
		tx.Rollback()
		return false, err
	}

	rows, err := tx.Query(`UPDATE balance SET balance = $1 WHERE user_id = $2`, sum+order.Accrual, userID)

	if err != nil || rows.Err() != nil {
		tx.Rollback()
		return false, err
	}

	tx.Commit()

	return true, nil
}
