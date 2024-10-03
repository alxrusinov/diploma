package model

import (
	luhn "github.com/EClaesson/go-luhn"
)

type Order struct {
	Number  string  `json:"order"`
	Process Process `json:"status"`
	Accrual int     `json:"accrual"`
}

func (order *Order) ValidateNumber() (bool, error) {
	res, err := luhn.IsValid(order.Number)
	return res, err
}
