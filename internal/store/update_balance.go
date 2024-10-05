package store

import "github.com/alxrusinov/diploma/internal/mathfn"

func (store *Store) UpdateBalance(balance float64, userID string) error {
	_, err := store.db.Exec(updateBalanceQuery, mathfn.RoundFloat(balance, 5), userID)

	if err != nil {
		return err
	}

	return nil

}
