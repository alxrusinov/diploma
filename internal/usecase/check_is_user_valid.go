package usecase

import "github.com/alxrusinov/diploma/internal/model"

func (usecase *Usecase) CheckIsValidUser(user *model.User) (string, error) {
	userID, err := usecase.store.FindUserByLoginPassword(user)

	return userID, err
}
