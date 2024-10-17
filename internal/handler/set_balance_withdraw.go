package handler

import (
	"errors"
	"net/http"

	"github.com/alxrusinov/diploma/internal/customerrors"
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

	if ok, err := withdraw.IsValid(); !ok || err != nil {
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

	err = handler.usecase.SetWithdrawls(withdraw, token.UserID)

	if err != nil {
		noMoneyErr := new(customerrors.PaymentRequiredError)

		if errors.As(err, &noMoneyErr) {
			ctx.AbortWithStatus(http.StatusPaymentRequired)
			return
		}

		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)

}
