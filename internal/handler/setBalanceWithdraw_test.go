package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alxrusinov/diploma/internal/authenticate"
	"github.com/alxrusinov/diploma/internal/model"
	"github.com/alxrusinov/diploma/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSetBalanceWithdraw(t *testing.T) {
	gin.SetMode(gin.TestMode)

	validWithdraw := &model.Withdrawn{
		Order:       "123",
		Sum:         500,
		ProcessedAt: "",
	}

	inValidWithdraw := &model.Withdrawn{
		Order:       "clown",
		Sum:         500,
		ProcessedAt: "",
	}

	noMoneyWithdraw := &model.Withdrawn{
		Order:       "123",
		Sum:         1,
		ProcessedAt: "",
	}

	testuseCase := new(usecase.UseCaseMock)

	testuseCase.On("UploadOrder", mock.Anything).Return(&model.Order{
		Number:  "123",
		Accrual: 400,
	}, nil)

	authClient := authenticate.CreateAuth()

	testHandler := CreateHandler(testuseCase, "http://localhost:8080", authClient)

	router := gin.New()

	router.POST("/api/user/withdraw", testHandler.SetBalanceWithDraw)

	tests := []struct {
		name     string
		withdraw *model.Withdrawn
		send     *model.Order
		body     []byte
		code     int
	}{
		{
			name:     "#1 bad body",
			withdraw: validWithdraw,
			send:     &model.Order{},
			body:     []byte("go"),
			code:     http.StatusInternalServerError,
		},
		{
			name:     "#2 invalid withdraw",
			withdraw: inValidWithdraw,
			send:     &model.Order{},
			body:     []byte(""),
			code:     http.StatusUnprocessableEntity,
		},
		{
			name:     "#3 no money",
			withdraw: noMoneyWithdraw,
			send:     &model.Order{},
			body:     []byte(""),
			code:     http.StatusPaymentRequired,
		},
		{
			name:     "#4 succes",
			withdraw: validWithdraw,
			send:     &model.Order{},
			body:     []byte(""),
			code:     http.StatusOK,
		},
	}

	for _, tt := range tests {
		send, _ := json.Marshal(tt.withdraw)

		if len(tt.body) != 0 {
			send = tt.body
		}

		request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/user/withdraw", bytes.NewReader(send))

		w := httptest.NewRecorder()

		router.ServeHTTP(w, request)

		res := w.Result()

		defer res.Body.Close()

		assert.Equal(t, tt.code, res.StatusCode)

	}
}
