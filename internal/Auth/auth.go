package auth

import (
	"time"

	"github.com/alxrusinov/diploma/internal/model"
	"github.com/golang-jwt/jwt"
)

type Auth struct {
	Sault []byte
}

func (auth *Auth) GetToken(user *model.User) (*model.Token, error) {
	exp := time.Now().Add(time.Hour * 600).Unix()

	payload := jwt.MapClaims{
		"sub": user.Login,
		"exp": exp,
	}
	tk := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenString, err := tk.SignedString(auth.Sault)

	if err != nil {
		return nil, err
	}

	token := &model.Token{
		UserName: user.Login,
		Exp:      exp,
		Token:    tokenString,
	}

	return token, nil
}

func CreateAuth() *Auth {
	return &Auth{
		Sault: []byte("quod licet jovi, non licet bovi"),
	}
}
