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

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Token struct {
	Token    string `json:"token"`
	UserName string `json:"user_name"`
	Exp      int64  `json:"exp"`
}
