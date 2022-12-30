package main

import (
	"github.com/andyj29/wannabet/internal/application/common"
	"github.com/andyj29/wannabet/internal/infrastructure/httpserver"
	"github.com/andyj29/wannabet/internal/infrastructure/logger"
	"github.com/andyj29/wannabet/internal/infrastructure/oauth"
	"github.com/andyj29/wannabet/internal/presentation/offer"
	"github.com/go-chi/chi/v5"
)

func main() {
	topLevelLogger := buildLogger("main", "main.log")
	server, err := httpserver.NewHTTPServer(httpserver.Config{
		Address: ":8080",
		Logger:  logger.InfraLogger,
		Router:  chi.NewRouter(),
	})
	if err != nil {
		topLevelLogger.Fatalf("Failed to initialize new http server")
	}

	cmdDispatcher := common.NewInMemoryDispatcher()

	offerController := offer.Controller{
		Dispatcher:           cmdDispatcher,
		GetRequestingAccount: oauth.ParseSub,
	}

	server.RegisterHandler("Post", "/api/create-offer", offerController.CreateOffer)
	server.RegisterHandler("Post", "/api/place-bet", offerController.PlaceBet)
}

func buildLogger(layer, logFile string) common.Logger {
	return logger.NewZapLogger(logger.Config{
		ServiceName: "Sport Betting",
		ServiceHost: "localhost",
		Layer:       layer,
		LogFileName: logFile,
	})
}
