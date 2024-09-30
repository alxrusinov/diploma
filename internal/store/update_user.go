package store

import "github.com/alxrusinov/diploma/internal/model"

func (store *Store) UpdateUser(token *model.Token) (*model.Token, error) {
	query := `UPDATE users SET token = $1 WHERE login = $2`

	tokenRows, err := store.db.Query(query, token.Token, token.UserName)

	if err != nil || tokenRows.Err() != nil {
		return nil, err
	}

	return token, nil
}
