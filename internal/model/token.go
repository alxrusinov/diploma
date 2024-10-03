package model

import "time"

type Token struct {
	UserID   string `json:"id"`
	Token    string `json:"token"`
	UserName string `json:"user_name"`
	Exp      int64  `json:"exp"`
}

func (token *Token) IsExpired() bool {
	return token.Exp < time.Now().Unix()
}
