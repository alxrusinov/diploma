package handler

import (
	"net/http"

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

	token, err := handler.AuthClient.GetToken(NewUser)

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	_, err = handler.store.UpdateUser(token)

	ctx.SetCookie(TokenCookie, token.Token, int(token.Exp), "/", "localhost", false, true)
	ctx.Status(http.StatusOK)
}
