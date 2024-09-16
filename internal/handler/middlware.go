package handler

import "github.com/gin-gonic/gin"

type Middleware struct {
}

func (middleware Middleware) CheckAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
