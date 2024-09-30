package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Middleware struct {
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

		if token.IsExpired() {
			ctx.SetCookie(TokenCookie, "", -1, "/", "localhost", false, true)
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Next()

	}
}
