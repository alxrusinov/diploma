package store

import "github.com/alxrusinov/diploma/internal/model"

func (store *Store) UpdateUser(token *model.Token) (*model.Token, error) {
	tokenRows, err := store.db.Query(updateUserTokenQuery, token.Token, token.UserName)

	if err != nil || tokenRows.Err() != nil {
		return nil, err
	}

	return token, nil
}
