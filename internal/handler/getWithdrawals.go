package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *Handler) GetWithdrawals(ctx *gin.Context) {
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

	balances, err := handler.useCase.GetWithdrawls(token.UserName)

	if len(balances) == 0 {
		ctx.AbortWithStatus(http.StatusNoContent)
		return
	}

	ctx.JSON(http.StatusOK, balances)
}
