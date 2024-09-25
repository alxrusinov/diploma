package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alxrusinov/diploma/internal/authenticate"
	"github.com/alxrusinov/diploma/internal/model"
	"github.com/alxrusinov/diploma/internal/useCase"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)

	trueUser := &model.User{
		Login:    "Ivan",
		Password: "1234",
	}

	badLoginUser := &model.User{
		Login:    "",
		Password: "1234",
	}

	badPasswordUser := &model.User{
		Login:    "Ivan",
		Password: "",
	}

	notValidUser := &model.User{
		Login:    "Petr",
		Password: "1234",
	}

	testUseCase := new(useCase.UseCaseMock)

	testUseCase.On("CheckUserExists", mock.Anything).Return(true, nil)

	testUseCase.On("CheckIsValidUser", notValidUser).Return(false, nil)
	testUseCase.On("CheckIsValidUser", trueUser).Return(true, nil)

	testUseCase.On("UpdateUser", mock.Anything).Return(&model.Token{
		UserName: trueUser.Login,
		Exp:      60 * 60 * 24,
		Token:    "123.456.789",
	}, nil)

	authClient := authenticate.CreateAuth()

	testHandler := CreateHandler(testUseCase, "http://localhost:8080", authClient)

	router := gin.New()

	router.POST("/api/user/login", testHandler.Login)

	tests := []struct {
		name   string
		body   []byte
		user   *model.User
		code   int
		cookie bool
	}{
		{
			name:   "#1 bad request",
			body:   []byte("123"),
			user:   nil,
			code:   http.StatusBadRequest,
			cookie: false,
		},
		{
			name:   "#2 empty login",
			body:   []byte(""),
			user:   badLoginUser,
			code:   http.StatusBadRequest,
			cookie: false,
		},
		{
			name:   "#3 empty password",
			body:   []byte(""),
			user:   badPasswordUser,
			code:   http.StatusBadRequest,
			cookie: false,
		},
		{
			name:   "#4 user is not valid",
			body:   []byte(""),
			user:   notValidUser,
			code:   http.StatusUnauthorized,
			cookie: false,
		},
		{
			name:   "#6 success",
			body:   []byte(""),
			user:   trueUser,
			code:   http.StatusOK,
			cookie: true,
		},
	}

	for _, tt := range tests {
		send, _ := json.Marshal(tt.user)

		if len(tt.body) != 0 {
			send = tt.body
		}

		request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/user/login", bytes.NewReader(send))
		w := httptest.NewRecorder()

		router.ServeHTTP(w, request)

		res := w.Result()

		assert.Equal(t, tt.code, res.StatusCode)

		var isCookieSet bool
		for _, cookie := range res.Cookies() {
			if cookie.Name == TokenCookie {
				isCookieSet = true
				break
			}
		}

		assert.Equal(t, tt.cookie, isCookieSet)
	}
}
