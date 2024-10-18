package handler

import (
	"net/http"

	"github.com/alxrusinov/diploma/internal/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (handler *Handler) GetBalance(ctx *gin.Context) {
	tokenString, err := ctx.Cookie(TokenCookie)

	if err != nil {
		logger.Logger.Error("error parsing cookie", zap.Error(err))
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	token, err := handler.AuthClient.ParseToken(tokenString)

	if err != nil {
		logger.Logger.Error("error parsing token", zap.Error(err))
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	balance, err := handler.usecase.GetBalance(token.UserID)

	if err != nil {
		logger.Logger.Error("error getting balance", zap.Error(err))
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	balance.Round()

	ctx.JSON(http.StatusOK, balance)
}
