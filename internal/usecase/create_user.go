package usecase

import (
	"github.com/alxrusinov/diploma/internal/model"
)

func (useCase *Usecase) CreateUser(user *model.User) (string, error) {
	userID, err := useCase.store.CreateUser(user)

	if err != nil {
		return "", err
	}

	return userID, nil
}
