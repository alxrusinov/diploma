package store

import (
	"context"

	"github.com/alxrusinov/diploma/internal/model"
)

func (store *Store) UpdateOrder(ctx context.Context, order *model.Order) error {

	tx, err := store.db.Begin()

	if err != nil {
		return err
	}

	select {
	case <-ctx.Done():
		tx.Rollback()
		return ctx.Err()
	default:
		if order.Process == model.Registered {
			order.Process = model.New
		}
		_, err := tx.Exec(updateOrderQuery, order.Process, order.Accrual, order.Number)

		if err != nil {
			tx.Rollback()
			return err
		}

		var balance float64

		err = tx.QueryRow(selectBalanceQuery, order.UserID).Scan(&balance)

		if err != nil {
			tx.Rollback()
			return err
		}

		_, err = tx.Exec(updateBalanceQuery, balance+order.Accrual, order.UserID)

		if err != nil {
			tx.Rollback()
			return err
		}

		tx.Commit()

		return nil

	}

}
