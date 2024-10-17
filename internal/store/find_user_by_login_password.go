package store

import (
	"errors"
	"io"

	"github.com/alxrusinov/diploma/internal/model"
)

func (store *Store) FindUserByLoginPassword(user *model.User) (string, error) {
	row := store.db.QueryRow(selectUserByLoiginPasswordQuery, user.Login, user.Password)

	var login string

	err := row.Scan(&login)

	if err != nil && !errors.Is(err, io.EOF) {
		return "", err
	}

	if login == "" {
		return "", nil
	}

	return login, nil
}
