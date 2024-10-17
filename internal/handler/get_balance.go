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

	balance, err := handler.usecase.GetBalance(token.UserID)

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	balance.Round()

	ctx.JSON(http.StatusOK, balance)
}
