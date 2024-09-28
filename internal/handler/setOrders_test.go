package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alxrusinov/diploma/internal/authenticate"
	"github.com/alxrusinov/diploma/internal/customerrors"
	"github.com/alxrusinov/diploma/internal/model"
	"github.com/alxrusinov/diploma/internal/use"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSetOrders(t *testing.T) {
	gin.SetMode(gin.TestMode)

	validOrder := &model.Order{
		Number: "123",
	}

	duplicateOwnerOrder := &model.Order{
		Number: "111",
	}

	duplicateUserOrder := &model.Order{
		Number: "222",
	}

	anotherErrorOrder := &model.Order{
		Number: "000",
	}

	badBody := &model.Order{
		Number: "999",
	}

	testuseCase := new(use.UsecaseMock)

	testuseCase.On("UploadOrder", validOrder).Return(validOrder, nil)

	testuseCase.On("UploadOrder", duplicateOwnerOrder).Return(duplicateOwnerOrder, new(customerrors.DuplicateOwnerOrderError))

	testuseCase.On("UploadOrder", duplicateUserOrder).Return(duplicateUserOrder, new(customerrors.DuplicateUserOrderError))

	testuseCase.On("UploadOrder", anotherErrorOrder).Return(anotherErrorOrder, errors.New("error"))

	authClient := authenticate.CreateAuth()

	testHandler := CreateHandler(testuseCase, "http://localhost:8080", authClient)

	router := gin.New()

	router.POST("/api/user/orders", testHandler.SetOrders)

	tests := []struct {
		name  string
		order *model.Order
		mock  *model.OrderMock
		code  int
	}{
		{
			name:  "#1 bad body",
			order: badBody,
			mock:  nil,
			code:  http.StatusBadRequest,
		},
		{
			name:  "#2 duplicateOwnerOrder",
			order: duplicateOwnerOrder,
			mock:  nil,
			code:  http.StatusOK,
		},
		{
			name:  "#3 duplicateUserOrder",
			order: duplicateUserOrder,
			mock:  nil,
			code:  http.StatusConflict,
		},
		{
			name:  "#4 anotherErrorOrder",
			order: anotherErrorOrder,
			mock:  nil,
			code:  http.StatusInternalServerError,
		},
		{
			name:  "#5 success",
			order: validOrder,
			mock:  nil,
			code:  http.StatusAccepted,
		},
	}

	for _, tt := range tests {
		send, _ := json.Marshal(tt.order.Number)

		if tt.order.Number == badBody.Number {
			send, _ = json.Marshal(tt.order)
		}

		request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/user/orders", bytes.NewReader(send))
		w := httptest.NewRecorder()

		router.ServeHTTP(w, request)

		res := w.Result()

		defer res.Body.Close()

		assert.Equal(t, tt.code, res.StatusCode)

	}

}
