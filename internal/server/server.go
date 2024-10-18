package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

type Server struct {
	mux        *gin.Engine
	handler    Handler
	runAddress string
	server     *http.Server
}

type Handler interface {
	GetBalance(ctx *gin.Context)
	GetOrders(ctx *gin.Context)
	GetWithdrawals(ctx *gin.Context)
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
	SetOrders(ctx *gin.Context)
	SetBalanceWithDraw(ctx *gin.Context)
	CheckAuth() gin.HandlerFunc
}

func (server *Server) Run() error {
	if err := server.server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func (server *Server) Shutdown(ctx context.Context) error {
	<-ctx.Done()
	if err := server.server.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}

func NewServer(handler Handler, runAddress string) *Server {
	server := &Server{
		mux:        gin.New(),
		handler:    handler,
		runAddress: runAddress,
	}

	api := server.mux.Group("/api")

	api.POST("/user/register", server.handler.Register)

	api.POST("/user/login", server.handler.Login)

	userAPI := api.Group("/user")

	userAPI.Use(server.handler.CheckAuth())

	userAPI.POST("/orders", server.handler.SetOrders)

	userAPI.GET("/orders", server.handler.GetOrders)

	userAPI.GET("/balance", server.handler.GetBalance)

	userAPI.POST("/balance/withdraw", server.handler.SetBalanceWithDraw)

	userAPI.GET("/withdrawals", server.handler.GetWithdrawals)

	server.server = &http.Server{
		Addr:    server.runAddress,
		Handler: server.mux,
	}

	return server
}
