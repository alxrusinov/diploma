package model

type Process string

const (
	Registered Process = "REGISTERED"
	Invalid    Process = "INVALID"
	Processing Process = "PROCESSING"
	Processed  Process = "PROCESSED"
)

type Order struct {
	Number  string  `json:"order"`
	Process Process `json:"status"`
	Accrual int     `json:"accrual"`
}
