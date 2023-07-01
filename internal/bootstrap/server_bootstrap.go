package bootstrap

import (
	"context"
	"fmt"
	"net/http"

	"github.com/krobus00/analytics-service/internal/config"
	"github.com/krobus00/analytics-service/internal/infrastructure"
	"github.com/krobus00/analytics-service/internal/repository"
	httpServer "github.com/krobus00/analytics-service/internal/transport/http"
	"github.com/krobus00/analytics-service/internal/usecase"
	"github.com/krobus00/analytics-service/internal/utils"
)

func StartServer() {
	// init infra
	infrastructure.InitializeDBConn()
	db, err := infrastructure.DB.DB()
	utils.ContinueOrFatal(err)

	echo := infrastructure.NewHTTPServer()

	// init repository
	analyticRepo := repository.NewAnalyticRepository()
	err = analyticRepo.InjectDB(infrastructure.DB)
	utils.ContinueOrFatal(err)

	// init usecase
	analyticUC := usecase.NewAnalyticUsecase()
	err = analyticUC.InjectAnalyticRepo(analyticRepo)
	utils.ContinueOrFatal(err)
	err = analyticUC.InjectDB(infrastructure.DB)
	utils.ContinueOrFatal(err)

	// init transport
	analyticCtrl := httpServer.NewAnalyticController()
	err = analyticCtrl.InjectAnalyticUC(analyticUC)
	utils.ContinueOrFatal(err)

	httpDelivery := httpServer.NewDelivery()
	err = httpDelivery.InjectEchoServer(echo)
	utils.ContinueOrFatal(err)
	err = httpDelivery.InjectAnalityCtrl(analyticCtrl)
	utils.ContinueOrFatal(err)

	httpDelivery.InitRoutes()

	go func() {
		if err := echo.Start(fmt.Sprintf(":%s", config.PortHTTP())); err != nil && err != http.ErrServerClosed {
			utils.ContinueOrFatal(err)
		}
	}()

	wait := gracefulShutdown(context.Background(), config.GracefulShutdownTimeOut(), map[string]operation{
		"database connection": func(ctx context.Context) error {
			infrastructure.StopTickerCh <- true
			return db.Close()
		},
		"echo": func(ctx context.Context) error {
			return echo.Shutdown(ctx)
		},
	})

	<-wait
}
