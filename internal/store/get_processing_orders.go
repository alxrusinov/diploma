package store

import (
	"context"

	"github.com/alxrusinov/diploma/internal/model"
)

func (store *Store) GetProcessingOrder(ctx context.Context) ([]model.Order, error) {
	res := make([]model.Order, 0)

	rows, err := store.db.Query(getProcessingOrder)

	if err != nil {
		return res, err
	}

	defer rows.Close()

	for rows.Next() {
		order := new(model.Order)

		err := rows.Scan(&order.UserID, &order.Number, &order.Process, &order.Accrual)

		if err != nil {
			return res, err
		}

		order.Round()

		res = append(res, *order)
	}

	if err := rows.Err(); err != nil {
		return res, err
	}

	return res, nil

}
