package model

import (
	"errors"

	luhn "github.com/EClaesson/go-luhn"
	"github.com/alxrusinov/diploma/internal/logger"
	"github.com/alxrusinov/diploma/internal/mathfn"
	"go.uber.org/zap"
)

type Order struct {
	Number  string  `json:"order"`
	Process Process `json:"status"`
	Accrual float64 `json:"accrual"`
	UserID  string  `json:"user_id,omitempty"`
}

func (order *Order) ValidateNumber() (bool, error) {
	if len(order.Number) == 0 {
		logger.Logger.Fatal("len number == 0")
		return false, errors.New("empty order number")
	}

	res, err := luhn.IsValid(order.Number)

	if err != nil {
		logger.Logger.Fatal("validate number err", zap.Any("err", err), zap.Bool("isValid", res))

	}

	return res, err
}

func (order *Order) Round() {
	order.Accrual = mathfn.RoundFloat(order.Accrual, 5)
}
