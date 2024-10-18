package app

import (
	"context"

	"github.com/alxrusinov/diploma/internal/authenticate"
	"github.com/alxrusinov/diploma/internal/client"
	"github.com/alxrusinov/diploma/internal/config"
	"github.com/alxrusinov/diploma/internal/handler"
	"github.com/alxrusinov/diploma/internal/logger"
	"github.com/alxrusinov/diploma/internal/migrator"
	"github.com/alxrusinov/diploma/internal/model"
	"github.com/alxrusinov/diploma/internal/orderclient"
	"github.com/alxrusinov/diploma/internal/server"
	"github.com/alxrusinov/diploma/internal/store"
	"github.com/alxrusinov/diploma/internal/usecase"
	"go.uber.org/zap"
)

type App struct {
	Config Config
}

type Config interface {
	Init()
	GetDatabaseURI() string
	GetAccrualSystemAddress() string
	GetRunAddress() string
	GetMigrationsDir() string
}

func (app *App) Run(ctx context.Context) chan error {
	errChan := make(chan error)
	app.Config.Init()

	err := logger.InitLogger()

	if err != nil {
		errChan <- err
	}

	migratorInst := migrator.NewMigrator()

	store := store.NewStore(app.Config.GetDatabaseURI(), migratorInst)

	authClient := authenticate.NewAuth()

	serviceClient := client.NewClient(app.Config.GetAccrualSystemAddress(), config.ClientTimeout)

	updateOrderClient := orderclient.NewOrderClient(store)

	uc := usecase.NewUsecase(store, serviceClient)

	router := handler.NewHandler(uc, app.Config.GetAccrualSystemAddress(), authClient)
	server := server.NewServer(router, app.Config.GetRunAddress())

	err = store.RunMigration()

	if err != nil {
		errChan <- err
	}

	go func(errChan chan error) {
		if err = server.Run(); err != nil {
			errChan <- err
		}
	}(errChan)

	go func(ctx context.Context, errChan chan error) {
		<-ctx.Done()
		if err = server.Shutdown(ctx); err != nil {
			errChan <- err
		}
	}(ctx, errChan)

	go func(ctx context.Context) {
		orderChan := updateOrderClient.GetProcessingOrder(ctx)

		orderInfoChan := serviceClient.GetOrderInfo(ctx, orderChan)

		updateOrderClient.UpdateOrder(ctx, orderInfoChan)

	}(ctx)

	ord := &model.Order{
		Number: string([]byte(`12345678902`)),
	}

	re, er := ord.ValidateNumber()

	logger.Logger.Error("ORDER", zap.Bool("RESULT", re), zap.Error(er))

	return errChan
}

func NewApp() *App {
	var appConfig = config.NewConfig()
	return &App{Config: appConfig}
}
