package handler

import (
	"net/http"
	"time"

	"github.com/alxrusinov/diploma/internal/auth"
	"github.com/alxrusinov/diploma/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *Handler) Register(ctx *gin.Context) {
	var NewUser = new(model.User)

	err := ctx.ShouldBindJSON(NewUser)

	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if NewUser.Login == "" || NewUser.Password == "" {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	isUserExist, err := handler.store.CheckUserExists(NewUser)

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if isUserExist {
		ctx.AbortWithStatus(http.StatusConflict)
		return
	}

	err = handler.store.CreateUser(NewUser)

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	tokeString, err := auth.GetToken(NewUser)

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	newToken := &model.Token{
		UserName: NewUser.Login,
		Exp:      time.Now().Add(time.Hour * 600).Unix(),
		Token:    tokeString,
	}

	token, err := handler.store.UpdateUser(newToken)

	ctx.SetCookie(TokenCookie, token.Token, int(token.Exp), "/", "localhost", false, true)
	ctx.Status(http.StatusOK)
}
