package usecase

import (
	"github.com/alxrusinov/diploma/internal/model"
)

func (usecase *Usecase) CheckUserExists(user *model.User) (bool, error) {
	ok, err := usecase.store.FindUserByLogin(user)
	return ok, err
}
