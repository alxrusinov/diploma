package usecase

import "github.com/alxrusinov/diploma/internal/model"

func (usecase *Usecase) UpdateUser(token *model.Token) (*model.Token, error) {
	resToken, err := usecase.store.UpdateUser(token)

	return resToken, err
}
