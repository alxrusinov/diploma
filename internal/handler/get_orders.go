package handler

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/alxrusinov/diploma/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *Handler) GetOrders(ctx *gin.Context) {
	tokenString, err := ctx.Cookie(TokenCookie)

	var defaultResp []model.Order

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, defaultResp)
		return
	}

	token, err := handler.AuthClient.ParseToken(tokenString)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, defaultResp)
		return
	}

	orders, err := handler.usecase.GetOrders(token.UserName)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNoContent, defaultResp)
			return
		}

		ctx.JSON(http.StatusInternalServerError, defaultResp)
		return
	}

	if len(orders) == 0 {
		fmt.Printf("%#v\n", orders)
		ctx.JSON(http.StatusNoContent, orders)
		return
	}

	ctx.JSON(http.StatusOK, orders)
}
