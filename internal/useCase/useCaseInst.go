package useCase

import (
	"github.com/alxrusinov/diploma/internal/model"
	"github.com/alxrusinov/diploma/internal/store"
)

type UseCaseInst struct {
	store store.Store
}

func (useCase *UseCaseInst) CheckUserExists(user *model.User) (bool, error) {
	return false, nil
}

func (useCase *UseCaseInst) CreateUser(user *model.User) error {
	return nil
}

func (useCase *UseCaseInst) UpdateUser(token *model.Token) (*model.Token, error) {
	return &model.Token{}, nil
}

func (useCase *UseCaseInst) CheckIsValidUser(user *model.User) (bool, error) {
	return true, nil
}

func (useCase *UseCaseInst) UploadOrder(order *model.Order) (*model.Order, error) {
	return &model.Order{}, nil
}

func (useCase *UseCaseInst) GetOrders(login string) ([]model.Order, error) {
	var res []model.Order
	return res, nil
}

func CreateUseCase(store store.Store) UseCase {
	return &UseCaseInst{store}
}
