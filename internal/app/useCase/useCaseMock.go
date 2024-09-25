package useCase

import (
	"github.com/alxrusinov/diploma/internal/model"
	"github.com/stretchr/testify/mock"
)

type UseCaseMock struct {
	mock.Mock
}

func (m *UseCaseMock) CheckUserExists(user *model.User) (bool, error) {
	args := m.Called(user)

	return args.Bool(0), args.Error(1)
}

func (m *UseCaseMock) CreateUser(user *model.User) error {
	args := m.Called(user)

	return args.Error(0)
}

func (m *UseCaseMock) UpdateUser(token *model.Token) (*model.Token, error) {
	args := m.Called(token)

	return args.Get(0).(*model.Token), args.Error(1)
}

func (m *UseCaseMock) CheckIsValidUser(user *model.User) (bool, error) {
	args := m.Called(user)

	return args.Bool(0), args.Error(1)
}

func (m *UseCaseMock) UploadOrder(order *model.Order) (*model.Order, error) {
	args := m.Called(order)

	return args.Get(0).(*model.Order), args.Error(1)
}
