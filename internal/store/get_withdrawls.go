package store

import (
	"github.com/alxrusinov/diploma/internal/customerrors"
	"github.com/alxrusinov/diploma/internal/model"
)

func (store *Store) GetWithdrawls(userID string) ([]model.Withdrawn, error) {
	res := make([]model.Withdrawn, 0)

	rows, err := store.db.Query(getWithdrawlsQuery, userID)

	if err != nil {
		return res, err
	}

	defer rows.Close()

	for rows.Next() {
		item := new(model.Withdrawn)

		err = rows.Scan(&item.Order, &item.Sum, &item.ProcessedAt)

		if err != nil {
			return res, err
		}
		item.Round()

		res = append(res, *item)
	}

	if err := rows.Err(); err != nil {
		return res, &customerrors.NoContentError{
			Err: err,
		}
	}

	return res, nil
}
