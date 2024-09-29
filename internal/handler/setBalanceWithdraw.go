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

	sendOrder := &model.Order{
		Number: withdraw.Order,
	}

	order, err := handler.usecase.UploadOrder(sendOrder, token.UserName)

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
