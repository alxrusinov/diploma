package model

import (
	luhn "github.com/EClaesson/go-luhn"
	"github.com/alxrusinov/diploma/internal/mathfn"
)

type Order struct {
	Number  string  `json:"order"`
	Process Process `json:"status"`
	Accrual float64 `json:"accrual"`
	UserID  string  `json:"user_id,omitempty"`
}

func (order *Order) ValidateNumber() (bool, error) {
	res, err := luhn.IsValid(order.Number)
	return res, err
}

func (order *Order) Round() {
	order.Accrual = mathfn.RoundFloat(order.Accrual, 5)
}
