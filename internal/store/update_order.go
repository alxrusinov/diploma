package store

import (
	"context"

	"github.com/alxrusinov/diploma/internal/model"
)

func (store *Store) UpdateOrder(ctx context.Context, userID string, orderCh <-chan *model.Order) (<-chan struct{}, <-chan error) {

	errCh := make(chan error)
	doneCh := make(chan struct{})

	go func() {
		tx, err := store.db.Begin()

		if err != nil {
			close(errCh)
			return
		}

		select {
		case <-ctx.Done():
			tx.Rollback()
			return
		case order := <-orderCh:
			if order.Process == model.Registered {
				order.Process = model.New
			}
			_, err := tx.Exec(updateOrderQuery, order.Process, order.Accrual, order.Number, userID)

			if err != nil {
				tx.Rollback()
				close(errCh)
				return
			}

			var balance float64

			err = tx.QueryRow(selectBalanceQuery, userID).Scan(&balance)

			if err != nil {
				tx.Rollback()
				close(errCh)
				return
			}

			_, err = tx.Exec(updateBalanceQuery, balance+order.Accrual, userID)

			if err != nil {
				tx.Rollback()
				close(errCh)
				return
			}

			if order.Process == model.Processed || order.Process == model.Invalid {
				tx.Commit()
				close(doneCh)
				return
			}

			tx.Commit()

		}
	}()

	return doneCh, errCh

}
