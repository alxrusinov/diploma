package handler

import (
	"net/http"

	"github.com/alxrusinov/diploma/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *Handler) Login(ctx *gin.Context) {
	var User = new(model.User)

	err := ctx.ShouldBindJSON(User)

	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if User.Login == "" || User.Password == "" {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	isValid, err := handler.useCase.CheckIsValidUser(User)

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if !isValid {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, err := handler.AuthClient.GetToken(User)

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	_, err = handler.useCase.UpdateUser(token)

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.SetCookie(TokenCookie, token.Token, int(token.Exp), "/", "localhost", false, true)
	ctx.Status(http.StatusOK)
}
