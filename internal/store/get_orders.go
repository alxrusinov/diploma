package store

import "github.com/alxrusinov/diploma/internal/model"

func (store *Store) GetOrders(login string) ([]model.OrderResponse, error) {
	res := make([]model.OrderResponse, 0)
	var userID string

	err := store.db.QueryRow(selectUserByLoginQuery, login).Scan(&userID)

	if err != nil {
		return res, err
	}

	rows, err := store.db.Query(selectOrdersQuery, userID)

	if err != nil {
		return res, err
	}

	defer rows.Close()

	for rows.Next() {
		item := new(model.OrderResponse)
		err = rows.Scan(&item.Number, &item.Status, &item.Accrual, &item.UploadedAt)

		if err != nil {
			return res, err
		}

		res = append(res, *item)

	}

	if rows.Err() != nil {
		return res, err
	}

	return res, nil
}
