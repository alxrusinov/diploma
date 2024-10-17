package model

import "github.com/stretchr/testify/mock"

type OrderMock struct {
	mock.Mock
}

func (m *OrderMock) ValidateNumber() bool {
	args := m.Called()

	return args.Bool(0)
}
