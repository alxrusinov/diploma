package authenticate

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

type CustomClaims struct {
	Exp int64  `json:"exp"`
	Sub string `json:"sub"`
	ID  string `json:"id"`
	jwt.StandardClaims
}

func (auth *Auth) GetToken(user *model.User) (*model.Token, error) {
	exp := time.Now().Add(time.Hour * 600).Unix()

	payload := jwt.MapClaims{
		"sub": user.Login,
		"exp": exp,
		"id":  user.ID,
	}
	tk := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenString, err := tk.SignedString(auth.Sault)

	if err != nil {
		return nil, err
	}

	token := &model.Token{
		UserID:   user.ID,
		UserName: user.Login,
		Exp:      exp,
		Token:    tokenString,
	}

	return token, nil
}

func (auth *Auth) ParseToken(tokenString string) (*model.Token, error) {
	var claims *CustomClaims

	parsed, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return auth.Sault, nil
	})

	if err != nil {
		return nil, err
	}

	if val, ok := parsed.Claims.(*CustomClaims); ok {
		claims = val
	} else {
		return nil, errors.New("claims parsing error")
	}

	token := &model.Token{
		UserID:   claims.ID,
		UserName: claims.Sub,
		Exp:      claims.Exp,
		Token:    tokenString,
	}

	return token, nil
}

func NewAuth() *Auth {
	return &Auth{
		Sault: []byte("quod licet jovi, non licet bovi"),
	}
}
