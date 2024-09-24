package handler

import (
	"io"
	"net/http"

	"github.com/alxrusinov/diploma/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *Handler) SetOrders(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	number := string(body)

	order := &model.Order{
		Number: number,
	}

	if !order.ValidateNumber() {
		ctx.AbortWithStatus(http.StatusUnprocessableEntity)
		return
	}

}
