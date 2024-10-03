package model

import (
	"errors"
	"strconv"

	luhn "github.com/EClaesson/go-luhn"
)

type Withdrawn struct {
	Order       string `json:"order"`
	Sum         int    `json:"sum"`
	ProcessedAt string `json:"processed_at"`
}

func (w *Withdrawn) IsWithdrawAvailable(sum int) bool {
	return w.Sum >= sum
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
