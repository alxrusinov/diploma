package auth

import (
	"errors"
	"fmt"
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

func (auth *Auth) ParseToken(tokenString string) (*model.Token, error) {
	parsed, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return auth.Sault, nil
	})

	if err != nil {
		return nil, err
	}

	var claims jwt.MapClaims

	if val, ok := parsed.Claims.(jwt.MapClaims); ok {
		claims = val
	} else {
		return nil, errors.New("claims parsing error")
	}

	token := new(model.Token)

	if userName, ok := claims["sub"].(string); ok {
		token.UserName = userName
	} else {
		return nil, errors.New("claims parsing error")
	}

	if exp, ok := claims["exp"].(int64); ok {
		token.Exp = exp
	} else {
		return nil, errors.New("claims parsing error")
	}

	token.Token = tokenString

	return token, nil
}

func CreateAuth() *Auth {
	return &Auth{
		Sault: []byte("quod licet jovi, non licet bovi"),
	}
}
