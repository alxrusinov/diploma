package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/alxrusinov/diploma/internal/authenticate"
	"github.com/alxrusinov/diploma/internal/model"
	"github.com/alxrusinov/diploma/internal/useCase"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetOrders(t *testing.T) {
	gin.SetMode(gin.TestMode)

	validLogin := "111"

	emptyListLogin := "222"

	errorLogin := "333"

	orderList := []model.OrderResponse{
		{
			Number:     "123",
			Status:     model.New,
			Accrual:    "123",
			UploadedAt: "2023-12-12",
		},
	}

	emptyOrderList := make([]model.OrderResponse, 0)

	testuseCase := new(useCase.UseCaseMock)

	testuseCase.On("GetOrders", validLogin).Return(orderList, nil)

	testuseCase.On("GetOrders", emptyListLogin).Return(emptyOrderList, nil)

	testuseCase.On("GetOrders", errorLogin).Return(emptyOrderList, errors.New("any error"))

	authClient := authenticate.CreateAuth()

	testHandler := CreateHandler(testuseCase, "http://localhost:8080", authClient)

	router := gin.New()

	router.GET("/api/user/orders", testHandler.GetOrders)

	tests := []struct {
		name  string
		login string
		code  int
	}{
		{
			name:  "#1 empty list",
			login: emptyListLogin,
			code:  http.StatusNoContent,
		},
		{
			name:  "#2 any error",
			login: errorLogin,
			code:  http.StatusInternalServerError,
		},
		{
			name:  "#1 success",
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

		request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/user/orders", nil)

		w := httptest.NewRecorder()

		request.AddCookie(&cookie)

		router.ServeHTTP(w, request)

		res := w.Result()

		assert.Equal(t, tt.code, res.StatusCode)

	}
}
