package usecase

import (
	"github.com/alxrusinov/diploma/internal/model"
	"github.com/stretchr/testify/mock"
)

type UsecaseMock struct {
	mock.Mock
}

func (m *UsecaseMock) CheckUserExists(user *model.User) (bool, error) {
	args := m.Called(user)

	return args.Bool(0), args.Error(1)
}

func (m *UsecaseMock) CreateUser(user *model.User) error {
	args := m.Called(user)

	return args.Error(0)
}

func (m *UsecaseMock) UpdateUser(token *model.Token) (*model.Token, error) {
	args := m.Called(token)

	return args.Get(0).(*model.Token), args.Error(1)
}

func (m *UsecaseMock) CheckIsValidUser(user *model.User) (bool, error) {
	args := m.Called(user)

	return args.Bool(0), args.Error(1)
}

func (m *UsecaseMock) UploadOrder(order *model.Order) (*model.Order, error) {
	args := m.Called(order)

	return args.Get(0).(*model.Order), args.Error(1)
}

func (m *UsecaseMock) GetOrders(login string) ([]model.OrderResponse, error) {
	args := m.Called(login)

	return args.Get(0).([]model.OrderResponse), args.Error(1)
}

func (m *UsecaseMock) GetBalance(login string) (*model.Balance, error) {
	args := m.Called(login)

	return args.Get(0).(*model.Balance), args.Error(1)
}

func (m *UsecaseMock) GetWithdrawls(login string) ([]model.Balance, error) {
	args := m.Called(login)

	return args.Get(0).([]model.Balance), args.Error(1)
}
