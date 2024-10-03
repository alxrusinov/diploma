package store

import (
	"database/sql"
	"errors"

	"github.com/alxrusinov/diploma/internal/customerrors"
	"github.com/alxrusinov/diploma/internal/model"
)

func (store *Store) CheckOrder(order *model.Order) (string, error) {
	var userID string

	err := store.db.QueryRow(checkOrderQuery, order.Number).Scan(&userID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", new(customerrors.NoOrderError)
		}

		return "", err
	}

	return userID, nil
}
