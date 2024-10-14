package store

import (
	"context"

	"github.com/alxrusinov/diploma/internal/model"
)

func (store *Store) UpdateOrder(ctx context.Context, userID string, orderCh <-chan *model.Order, cancel context.CancelFunc) {

	select {
	case <-ctx.Done():
		return
	case order := <-orderCh:
		if order.Process == model.Registered {
			order.Process = model.New
		}
		_, err := store.db.Exec(updateOrderQuery, order.Process, order.Accrual, order.Number, userID)

		if err != nil || order.Process == model.Processed || order.Process == model.Invalid {
			cancel()
		}

	}

}
