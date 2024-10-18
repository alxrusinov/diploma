package handler

import "github.com/gin-gonic/gin"

func (handler *Handler) CheckAuth() gin.HandlerFunc {
	return handler.Middleware.CheckAuth(handler.AuthClient)
}
