package store

import "github.com/alxrusinov/diploma/internal/model"

func (store *Store) CreateUser(user *model.User) (bool, error) {
	tx, err := store.db.Begin()

	if err != nil {
		return false, err
	}
	queryUser := `INSERT INTO users (login, password, token)
				VALUES ($1, $2, $3)
				RETURNING id;`

	queryBalance := `INSERT INTO users (user_id, balance)
				VALUES ($1, $2);`

	var userID string

	err = tx.QueryRow(queryUser, user.Login, user.Password, "").Scan(&userID)

	if err != nil {
		tx.Rollback()
		return false, err
	}

	_, err = tx.Exec(queryBalance, userID, 0)

	if err != nil {
		tx.Rollback()
		return false, err
	}

	tx.Commit()

	return true, nil
}
