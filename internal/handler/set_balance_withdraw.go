package handler

import (
	"net/http"
	"time"

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

	balance, err := handler.usecase.GetBalance(token.UserID)

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	balanceInt := int(balance.Current)

	if !withdraw.IsWithdrawAvailable(balanceInt) {
		ctx.AbortWithStatus(http.StatusPaymentRequired)
		return
	}

	err = handler.usecase.UpdateBalance(balanceInt-withdraw.Sum, token.UserID)

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	withdraw.ProcessedAt = time.Now().Format(time.RFC3339)

	err = handler.usecase.SetWithdrawls(withdraw, token.UserID)

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)

}
