package handler

import (
	"errors"
	"io"
	"net/http"

	"github.com/alxrusinov/diploma/internal/customerrors"
	"github.com/alxrusinov/diploma/internal/logger"
	"github.com/alxrusinov/diploma/internal/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (handler *Handler) SetOrders(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)

	if err != nil {
		logger.Logger.Error("bad request of stting order", zap.Error(err))
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	parsedBody := string(body)

	order := &model.Order{
		Number: parsedBody,
	}

	isValid, err := order.ValidateNumber()

	if err != nil || !isValid {
		logger.Logger.Error("invalid order number", zap.Error(err))
		ctx.AbortWithStatus(http.StatusUnprocessableEntity)
		return
	}

	tokenString, err := ctx.Cookie(TokenCookie)

	if err != nil {
		logger.Logger.Error("invalid cookie", zap.Error(err))
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	token, err := handler.AuthClient.ParseToken(tokenString)

	if err != nil {
		logger.Logger.Error("invalid token", zap.Error(err))
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	err = handler.usecase.UploadOrder(order, token.UserID)

	if err != nil {
		duplicateOwnerError := new(customerrors.DuplicateOwnerOrderError)
		DuplicateUserOrderError := new(customerrors.DuplicateUserOrderError)

		if errors.As(err, &duplicateOwnerError) {
			logger.Logger.Error("user has alredy uploaded order", zap.Error(err), zap.String("UserID", order.UserID), zap.String("order number", order.Number))
			ctx.Status(http.StatusOK)
			return
		}

		if errors.As(err, &DuplicateUserOrderError) {
			logger.Logger.Error("another user has alredy uploaded order", zap.Error(err), zap.String("order number", order.Number))
			ctx.AbortWithStatus(http.StatusConflict)
			return
		}

		logger.Logger.Error("another uploading order error", zap.Error(err), zap.String("order number", order.Number))
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return

	}

	logger.Logger.Info("order is processing", zap.String("order number", order.Number))
	ctx.Status(http.StatusAccepted)

}
