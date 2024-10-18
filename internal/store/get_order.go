package store

import (
	"database/sql"
	"errors"

	"github.com/alxrusinov/diploma/internal/model"
)

func (store *Store) GetOrder(order *model.Order, userID string) (*model.Order, error) {
	resOrder := new(model.Order)

	err := store.db.QueryRow(getOrderQuery, order.Number, userID).Scan(&resOrder)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	resOrder.Round()

	return resOrder, nil
}
