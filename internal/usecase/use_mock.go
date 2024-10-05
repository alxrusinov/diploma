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

func (m *UsecaseMock) CreateUser(user *model.User) (string, error) {
	args := m.Called(user)

	return args.String(0), args.Error(1)
}

func (m *UsecaseMock) UpdateUser(token *model.Token) (*model.Token, error) {
	args := m.Called(token)

	return args.Get(0).(*model.Token), args.Error(1)
}

func (m *UsecaseMock) CheckIsValidUser(user *model.User) (string, error) {
	args := m.Called(user)

	return args.String(0), args.Error(1)
}

func (m *UsecaseMock) UploadOrder(order *model.Order, login string) error {
	args := m.Called(order, login)

	return args.Error(0)
}

func (m *UsecaseMock) GetOrders(login string) ([]model.OrderResponse, error) {
	args := m.Called(login)

	return args.Get(0).([]model.OrderResponse), args.Error(1)
}

func (m *UsecaseMock) GetBalance(login string) (*model.Balance, error) {
	args := m.Called(login)

	return args.Get(0).(*model.Balance), args.Error(1)
}

func (m *UsecaseMock) GetWithdrawls(login string) ([]model.Withdrawn, error) {
	args := m.Called(login)

	return args.Get(0).([]model.Withdrawn), args.Error(1)
}

func (m *UsecaseMock) GetOrder(order *model.Order, userID string) (*model.Order, error) {
	args := m.Called(order, userID)

	return args.Get(0).(*model.Order), args.Error(1)
}

func (m *UsecaseMock) UpdateBalance(balance float32, userID string) error {
	args := m.Called(balance, userID)

	return args.Error(0)
}

func (m *UsecaseMock) SetWithdrawls(withdrawn *model.Withdrawn, userID string) error {
	args := m.Called(withdrawn, userID)

	return args.Error(0)

}
