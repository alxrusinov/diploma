package model

import (
	"errors"
	"strconv"

	luhn "github.com/EClaesson/go-luhn"
	"github.com/alxrusinov/diploma/internal/mathfn"
)

type Withdrawn struct {
	Order       string  `json:"order"`
	Sum         float64 `json:"sum"`
	ProcessedAt string  `json:"processed_at"`
}

func (w *Withdrawn) IsWithdrawAvailable(balance float64) bool {
	return balance >= w.Sum
}

func (w *Withdrawn) IsValid() (bool, error) {
	res, err := luhn.IsValid(w.Order)
	return res, err
}

func (w *Withdrawn) OrderToNumber() (int, error) {
	order, err := strconv.Atoi(w.Order)

	if err != nil {
		return 0, errors.New("invalid order")
	}

	return order, nil
}

func (w *Withdrawn) Round() {
	w.Sum = mathfn.RoundFloat(w.Sum, 5)
}
