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

	body, err := io.ReadAll(ctx.Request.Body)

	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	parsedBody := string(body)

	order := &model.Order{
		Number: parsedBody,
		UserID: token.UserID,
	}

	isValid, err := order.ValidateNumber()

	if err != nil || !isValid {
		logger.Logger.Error("error order number", zap.Error(err))
		ctx.AbortWithStatus(http.StatusUnprocessableEntity)
		return
	}

	err = handler.usecase.UploadOrder(order, token.UserID)

	if err != nil {
		duplicateOwnerError := new(customerrors.DuplicateOwnerOrderError)
		duplicateUserOrderError := new(customerrors.DuplicateUserOrderError)

		if errors.As(err, &duplicateOwnerError) {
			logger.Logger.Info("user has alredy uploaded order", zap.Error(err), zap.String("UserID", order.UserID), zap.String("order number", order.Number))
			ctx.Status(http.StatusOK)
			return
		}

		if errors.As(err, &duplicateUserOrderError) {
			logger.Logger.Info("another user has alredy uploaded order", zap.Error(err), zap.String("order number", order.Number))
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
