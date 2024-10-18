package usecase

import (
	"github.com/alxrusinov/diploma/internal/model"
)

func (usecase *Usecase) CreateUser(user *model.User) (string, error) {
	userID, err := usecase.store.CreateUser(user)

	if err != nil {
		return "", err
	}

	return userID, nil
}
