package handler

import (
	"net/http"

	"github.com/alxrusinov/diploma/internal/model"
	"github.com/gin-gonic/gin"
)

type Middleware struct {
	usecase Usecase
}

func (middleware Middleware) CheckAuth(authClient Auth) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString, err := ctx.Cookie(TokenCookie)

		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token, err := authClient.ParseToken(tokenString)

		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		isUserExist, err := middleware.usecase.CheckUserExists(&model.User{
			ID:    token.UserID,
			Login: token.UserName,
		})

		if !isUserExist {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return

		}

		if token.IsExpired() {
			ctx.SetCookie(TokenCookie, "", -1, "/", "localhost", false, true)
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Next()

	}
}
