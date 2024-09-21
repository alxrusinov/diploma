package model

import "time"

type Token struct {
	Token    string `json:"token"`
	UserName string `json:"user_name"`
	Exp      int64  `json:"exp"`
}

func (token *Token) IsExpired() bool {
	return token.Exp < time.Now().Unix()
}
