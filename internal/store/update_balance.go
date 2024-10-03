package store

func (store *Store) UpdateBalance(balance int, userID string) error {
	_, err := store.db.Exec(updateBalanceQuery, balance, userID)

	if err != nil {
		return err
	}

	return nil

}
