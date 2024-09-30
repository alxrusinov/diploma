package usecase

import "github.com/alxrusinov/diploma/internal/model"

func (useCase *Usecase) CheckIsValidUser(user *model.User) (bool, error) {
	found, err := useCase.store.FindUserByLoginPassword(user)

	return found, err
}
