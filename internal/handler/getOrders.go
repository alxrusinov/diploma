package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *Handler) GetOrders(ctx *gin.Context) {
	tokenString, err := ctx.Cookie(TokenCookie)

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	token, err := handler.AuthClient.ParseToken(tokenString)

	if err != nil {
		fmt.Printf("#2 TWO - %#v\n", tokenString)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	orders, err := handler.useCase.GetOrders(token.UserName)

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if len(orders) == 0 {
		ctx.Status(http.StatusNoContent)
		return
	}

	ctx.JSON(http.StatusOK, orders)
}
