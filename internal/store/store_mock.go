package store

import (
	"github.com/alxrusinov/diploma/internal/model"
	"github.com/stretchr/testify/mock"
)

type DBStoreMock struct {
	mock.Mock
}

func (m *DBStoreMock) CheckUserExists(user *model.User) (bool, error) {
	args := m.Called(user)

	return args.Bool(0), args.Error(1)
}

func (m *DBStoreMock) CreateUser(user *model.User) error {
	args := m.Called(user)

	return args.Error(0)
}

func (m *DBStoreMock) UpdateUser(token *model.Token) (*model.Token, error) {
	args := m.Called(token)

	return args.Get(0).(*model.Token), args.Error(1)
}

func (m *DBStoreMock) CheckIsValidUser(user *model.User) (bool, error) {
	args := m.Called(user)

	return args.Bool(0), args.Error(1)
}

func (m *DBStoreMock) UploadOrder(order *model.Order, login string) error {
	args := m.Called(order)

	return args.Error(0)
}
