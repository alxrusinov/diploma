package store

import (
	"strconv"

	"github.com/alxrusinov/diploma/internal/model"
)

func (store *Store) UpdateUser(token *model.Token) (*model.Token, error) {
	numID, err := strconv.Atoi(token.UserID)

	if err != nil {
		return nil, err
	}

	tokenRows, err := store.db.Query(updateUserTokenQuery, token.Token, numID)

	if err != nil || tokenRows.Err() != nil {
		return nil, err
	}

	return token, nil
}
