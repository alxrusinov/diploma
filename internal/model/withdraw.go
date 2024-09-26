package model

import (
	"errors"
	"strconv"
)

type Withdrawn struct {
	Order       string `json:"order"`
	Sum         int    `json:"sum"`
	ProcessedAt string `json:"processed_at"`
}

func (w *Withdrawn) IsWithdrawAvailable(sum int) bool {
	return w.Sum >= sum
}

func (w *Withdrawn) IsValid() bool {
	return true
}

func (w *Withdrawn) OrderToNumber() (int, error) {
	order, err := strconv.Atoi(w.Order)

	if err != nil {
		return 0, errors.New("invalid order")
	}

	return order, nil
}
