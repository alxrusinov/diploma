package handler

import (
	"database/sql"
	"errors"
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
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	orders, err := handler.usecase.GetOrders(token.UserName)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNoContent, orders)
			return
		}

		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if len(orders) == 0 {
		ctx.JSON(http.StatusNoContent, orders)
		return
	}

	ctx.JSON(http.StatusOK, orders)
}
