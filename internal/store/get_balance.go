package store

import "github.com/alxrusinov/diploma/internal/model"

func (store *Store) GetBalance(userID string) (*model.Balance, error) {
	balance := new(model.Balance)

	err := store.db.QueryRow(selectBalanceQuery, userID).Scan(&balance.Current)

	if err != nil {
		return nil, err
	}

	return balance, nil
}
