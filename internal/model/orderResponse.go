package model

type OrderResponse struct {
	Number     string  `json:"number"`
	Status     Process `json:"status"`
	Accrual    string  `json:"accrual"`
	UploadedAt string  `json:"uploaded_at"`
}
