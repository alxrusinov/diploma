package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alxrusinov/diploma/internal/model"
	"github.com/alxrusinov/diploma/internal/store"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegister(t *testing.T) {
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

	existedUser := &model.User{
		Login:    "Petr",
		Password: "1234",
	}

	notCreatedUser := &model.User{
		Login:    "Modest",
		Password: "1234",
	}

	testStore := new(store.DBStoreMock)

	testStore.On("CheckUserExists", existedUser).Return(true, nil)
	testStore.On("CheckUserExists", trueUser).Return(false, nil)
	testStore.On("CheckUserExists", notCreatedUser).Return(false, nil)

	testStore.On("CreateUser", notCreatedUser).Return(errors.New("user was not created"))
	testStore.On("CreateUser", trueUser).Return(nil)

	testStore.On("UpdateUser", mock.Anything).Return(&model.Token{
		UserName: trueUser.Login,
		Exp:      60 * 60 * 24,
		Token:    "123.456.789",
	}, nil)

	testHandler := CreateHandler(testStore, "http://localhost:8080")

	router := gin.New()

	router.POST("/api/user/register", testHandler.Register)

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
			name:   "#4 user exists",
			body:   []byte(""),
			user:   existedUser,
			code:   http.StatusConflict,
			cookie: false,
		},
		{
			name:   "#5 user was not created",
			body:   []byte(""),
			user:   notCreatedUser,
			code:   http.StatusInternalServerError,
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

		request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/user/register", bytes.NewReader(send))
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
