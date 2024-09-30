package store

import (
	"errors"
	"io"

	"github.com/alxrusinov/diploma/internal/model"
)

func (store *Store) FindUserByLogin(user *model.User) (bool, error) {

	row := store.db.QueryRow(selectUserByLoginQuery, user.Login)

	var login string

	err := row.Scan(&login)

	if err != nil && !errors.Is(err, io.EOF) {
		return false, err
	}

	if login == "" {
		return false, nil
	}

	return true, nil
}
