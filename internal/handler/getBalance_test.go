package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/alxrusinov/diploma/internal/authenticate"
	"github.com/alxrusinov/diploma/internal/model"
	"github.com/alxrusinov/diploma/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetBalance(t *testing.T) {
	gin.SetMode(gin.TestMode)

	validLogin := "111"

	errorLogin := "333"

	testuseCase := new(usecase.UseCaseMock)

	testuseCase.On("GetBalance", validLogin).Return(&model.Balance{}, nil)

	testuseCase.On("GetBalance", errorLogin).Return(&model.Balance{}, errors.New("some errors"))

	authClient := authenticate.CreateAuth()

	testHandler := CreateHandler(testuseCase, "http://localhost:8080", authClient)

	router := gin.New()

	router.GET("/api/user/balance", testHandler.GetBalance)

	tests := []struct {
		name  string
		login string
		code  int
	}{
		{
			name:  "1# any error",
			login: errorLogin,
			code:  http.StatusInternalServerError,
		},
		{
			name:  "2# success",
			login: validLogin,
			code:  http.StatusOK,
		},
	}

	for _, tt := range tests {
		user := &model.User{
			Login: tt.login,
		}
		token, _ := testHandler.AuthClient.GetToken(user)

		cookie := http.Cookie{
			Name:    TokenCookie,
			Value:   token.Token,
			Expires: time.Now().Add(time.Hour * 24),
			Path:    "/",
		}

		request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/user/balance", nil)

		w := httptest.NewRecorder()

		request.AddCookie(&cookie)

		router.ServeHTTP(w, request)

		res := w.Result()

		defer res.Body.Close()

		assert.Equal(t, tt.code, res.StatusCode)
	}
}
