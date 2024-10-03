package usecase

import "github.com/alxrusinov/diploma/internal/model"

func (useCase *Usecase) CheckIsValidUser(user *model.User) (string, error) {
	userID, err := useCase.store.FindUserByLoginPassword(user)

	return userID, err
}
