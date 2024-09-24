package model

type Order struct {
	Number  string  `json:"order"`
	Process Process `json:"status"`
	Accrual int     `json:"accrual"`
}

func (order *Order) ValidateNumber() bool {
	return true
}
