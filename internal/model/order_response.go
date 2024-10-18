package model

import "github.com/alxrusinov/diploma/internal/mathfn"

type OrderResponse struct {
	Number     string  `json:"number"`
	Status     Process `json:"status"`
	Accrual    float64 `json:"accrual"`
	UploadedAt string  `json:"uploaded_at"`
	UserID     string  `json:"user_id,omitempty"`
}

func (o *OrderResponse) Round() {
	o.Accrual = mathfn.RoundFloat(o.Accrual, 5)
}
