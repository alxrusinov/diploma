package server

import (
	"github.com/alxrusinov/diploma/internal/handler"
	"github.com/gin-gonic/gin"
)

type Server struct {
	mux        *gin.Engine
	handler    *handler.Handler
	runAddress string
}

func (server *Server) Run() {
	server.mux.Run(server.runAddress)
}

func CreateServer(handler *handler.Handler, runAddress string) *Server {
	server := &Server{
		mux:        gin.New(),
		handler:    handler,
		runAddress: runAddress,
	}

	api := server.mux.Group("/api")

	api.POST("/user/register", server.handler.Register)

	api.POST("/user/login", server.handler.Login)

	userAPI := api.Group("/user")

	userAPI.Use(server.handler.Middleware.CheckAuth())

	userAPI.POST("/orders", server.handler.SetOrders)

	userAPI.GET("/orders", server.handler.GetOrders)

	userAPI.GET("/balance", server.handler.GetBalance)

	userAPI.POST("/balance/withdraw", server.handler.SetBalanceWithDraw)

	userAPI.GET("/withdrawals", server.handler.GetWithdrawals)

	return server
}
