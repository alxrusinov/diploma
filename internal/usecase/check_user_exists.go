package usecase

import (
	"github.com/alxrusinov/diploma/internal/model"
)

func (useCase *Usecase) CheckUserExists(user *model.User) (bool, error) {
	ok, err := useCase.store.FindUserByLogin(user)
	return ok, err
}
