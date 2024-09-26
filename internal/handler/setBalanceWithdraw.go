package handler

import (
	"net/http"

	"github.com/alxrusinov/diploma/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *Handler) SetBalanceWithDraw(ctx *gin.Context) {
	withdraw := new(model.Withdrawn)

	err := ctx.ShouldBindJSON(withdraw)

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if !withdraw.IsValid() {
		ctx.AbortWithStatus(http.StatusUnprocessableEntity)
		return
	}

	orderNumber, err := withdraw.OrderToNumber()

	if err != nil {
		ctx.AbortWithStatus(http.StatusUnprocessableEntity)
		return
	}

	sendOrder := &model.Order{
		Number: orderNumber,
	}

	order, err := handler.useCase.UploadOrder(sendOrder)

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if !withdraw.IsWithdrawAvailable(order.Accrual) {
		ctx.AbortWithStatus(http.StatusPaymentRequired)
		return
	}

	ctx.Status(http.StatusOK)

}
