package usecase

import (
	"errors"

	"github.com/alxrusinov/diploma/internal/model"
)

func (useCase *Usecase) CheckUserExists(user *model.User) (bool, error) {
	ok, err := useCase.store.FindUserByLogin(user)
	return ok, err
}

func (useCase *Usecase) CreateUser(user *model.User) error {
	ok, err := useCase.store.CreateUser(user)

	if err != nil {
		return err
	}

	if !ok {
		return errors.New("user was not created")
	}

	return nil
}
