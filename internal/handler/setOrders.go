package handler

import (
	"errors"
	"io"
	"net/http"

	"github.com/alxrusinov/diploma/internal/customerrors"
	"github.com/alxrusinov/diploma/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *Handler) SetOrders(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)

	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	parsedBody := string(body)

	order := &model.Order{
		Number: parsedBody,
	}

	isValid, err := order.ValidateNumber()

	if err != nil || !isValid {
		ctx.AbortWithStatus(http.StatusUnprocessableEntity)
		return
	}

	_, err = handler.usecase.UploadOrder(order)

	if err != nil {
		duplicateOwnerError := new(customerrors.DuplicateOwnerOrderError)
		DuplicateUserOrderError := new(customerrors.DuplicateUserOrderError)

		if errors.As(err, &duplicateOwnerError) {
			ctx.Status(http.StatusOK)
			return
		}

		if errors.As(err, &DuplicateUserOrderError) {
			ctx.AbortWithStatus(http.StatusConflict)
			return
		}

		ctx.AbortWithStatus(http.StatusInternalServerError)
		return

	}

	ctx.Status(http.StatusAccepted)

}
