package store

import "github.com/alxrusinov/diploma/internal/model"

const ()

func (store *Store) CreateUser(user *model.User) (bool, error) {
	tx, err := store.db.Begin()

	if err != nil {
		return false, err
	}

	var userID string

	err = tx.QueryRow(insertUserQuery, user.Login, user.Password, "").Scan(&userID)

	if err != nil {
		tx.Rollback()
		return false, err
	}

	_, err = tx.Exec(insertBalanceQuery, userID, 0)

	if err != nil {
		tx.Rollback()
		return false, err
	}

	tx.Commit()

	return true, nil
}
