package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *Handler) GetBalance(ctx *gin.Context) {
	tokenString, err := ctx.Cookie(TokenCookie)

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	token, err := handler.AuthClient.ParseToken(tokenString)

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	balance, err := handler.useCase.GetBalance(token.UserName)

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, balance)
}
