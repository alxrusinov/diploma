package handler

import (
	"fmt"
	"net/http"

	"github.com/alxrusinov/diploma/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *Handler) Register(ctx *gin.Context) {
	var NewUser = new(model.User)

	err := ctx.ShouldBindJSON(NewUser)

	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if NewUser.Login == "" || NewUser.Password == "" {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	isUserExist, err := handler.usecase.CheckUserExists(NewUser)

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if isUserExist {
		ctx.AbortWithStatus(http.StatusConflict)
		return
	}

	userID, err := handler.usecase.CreateUser(NewUser)

	if err != nil {
		fmt.Println("two")
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	NewUser.ID = userID

	token, err := handler.AuthClient.GetToken(NewUser)

	if err != nil {
		fmt.Println("three")
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	_, err = handler.usecase.UpdateUser(token)

	if err != nil {
		fmt.Println("four")
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.SetCookie(TokenCookie, token.Token, int(token.Exp), "/", "localhost", false, true)
	ctx.Status(http.StatusOK)
}
