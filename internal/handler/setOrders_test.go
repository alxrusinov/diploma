package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alxrusinov/diploma/internal/app/useCase"
	"github.com/alxrusinov/diploma/internal/auth"
	"github.com/alxrusinov/diploma/internal/customerrors"
	"github.com/alxrusinov/diploma/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSetOrders(t *testing.T) {
	gin.SetMode(gin.TestMode)

	validOrder := &model.Order{
		Number: 123,
	}

	duplicateOwnerOrder := &model.Order{
		Number: 111,
	}

	duplicateUserOrder := &model.Order{
		Number: 222,
	}

	anotherErrorOrder := &model.Order{
		Number: 000,
	}

	invalidNumberOrder := new(model.OrderMock)

	invalidNumberOrder.On("ValidateNumber").Return(false)

	testuseCase := new(useCase.UseCaseMock)

	testuseCase.On("UploadOrder", validOrder).Return(validOrder, nil)

	testuseCase.On("UploadOrder", duplicateOwnerOrder).Return(nil, new(customerrors.DuplicateOwnerOrderError))

	testuseCase.On("UploadOrder", duplicateUserOrder).Return(nil, new(customerrors.DuplicateUserOrderError))

	testuseCase.On("UploadOrder", anotherErrorOrder).Return(nil, errors.New("error"))

	authClient := auth.CreateAuth()

	testHandler := CreateHandler(testuseCase, "http://localhost:8080", authClient)

	router := gin.New()

	router.POST("/api/user/orders", testHandler.Register)

	tests := []struct {
		name  string
		body  []byte
		order *model.Order
		mock  *model.OrderMock
		code  int
	}{
		{
			name:  "#1 bad body",
			body:  []byte("clown"),
			order: nil,
			mock:  nil,
			code:  http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		send, _ := json.Marshal(tt.order)

		if len(tt.body) != 0 {
			send = tt.body
		}

		request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/user/orders", bytes.NewReader(send))
		w := httptest.NewRecorder()

		router.ServeHTTP(w, request)

		res := w.Result()

		assert.Equal(t, tt.code, res.StatusCode)

	}

}
