package usecase

import "github.com/alxrusinov/diploma/internal/model"

func (useCase *Usecase) UpdateUser(token *model.Token) (*model.Token, error) {
	resToken, err := useCase.store.UpdateUser(token)

	return resToken, err
}
