package main

import (
	"log"
	"os"

	"github.com/go-chi/chi/v5"

	"github.com/andyj29/wannabet/internal/api/controller"
	appAccount "github.com/andyj29/wannabet/internal/application/account"
	"github.com/andyj29/wannabet/internal/application/common"
	appOffer "github.com/andyj29/wannabet/internal/application/offer"
	"github.com/andyj29/wannabet/internal/domain/account"
	"github.com/andyj29/wannabet/internal/domain/offer"
	"github.com/andyj29/wannabet/internal/infrastructure/asyncbus"
	"github.com/andyj29/wannabet/internal/infrastructure/datastore"
	"github.com/andyj29/wannabet/internal/infrastructure/httpserver"
	"github.com/andyj29/wannabet/internal/infrastructure/logger"
	"github.com/andyj29/wannabet/internal/infrastructure/oauth"
)

func main() {
	server, err := httpserver.NewHTTPServer(httpserver.Config{
		Address: os.Getenv("SERVICE_ADDR"),
		Logger:  logger.InfraLogger,
		Router:  chi.NewRouter(),
	})
	if err != nil {
		log.Fatalf("Failed to initialize new HTTP server")
	}

	cmdDispatcher := common.NewInMemoryDispatcher()
	if err := buildAndRegisterCmdHandlers(cmdDispatcher); err != nil {
		log.Fatalf("Failed to build and register command handlers to dispatcher")
	}

	accountController := controller.AccountController{
		Dispatcher: cmdDispatcher,
	}

	offerController := controller.OfferController{
		Dispatcher:           cmdDispatcher,
		GetRequestingAccount: oauth.ParseSub,
	}

	server.RegisterHandler("Post", "/api/register-account", accountController.RegisterAccount)
	server.RegisterHandler("Post", "/api/create-offer", offerController.CreateOffer)
	server.RegisterHandler("Post", "/api/place-bet", offerController.PlaceBet)

}

func buildLogger(layer, logFile string) common.Logger {
	return logger.NewZapLogger(logger.Config{
		ServiceName: os.Getenv("SERVICE_NAME"),
		ServiceHost: os.Getenv("SERVICE_HOST"),
		Layer:       layer,
		LogFileName: logFile,
	})
}

func buildAndRegisterCmdHandlers(cmdDispatcher common.Dispatcher) error {
	kafkaProducer := asyncbus.NewProducer(os.Getenv("KAFKA_ADDR"))
	eventBus := asyncbus.NewAsyncEventBus(kafkaProducer)

	eventStore := datastore.NewEventStore(os.Getenv("EVENTSTORE_ADDR"))
	accountRepo := datastore.NewRepository[*account.Account](eventStore)
	offerRepo := datastore.NewRepository[*offer.Offer](eventStore)

	registerAccountHandler := appAccount.RegisterAccountHandler{
		Repo:     accountRepo,
		EventBus: eventBus,
	}
	if err := cmdDispatcher.RegisterHandler(appAccount.RegisterAccount{}, &registerAccountHandler); err != nil {
		return err
	}

	addFundsHandler := appAccount.AddFundsHandler{
		Repo:     accountRepo,
		EventBus: eventBus,
	}
	if err := cmdDispatcher.RegisterHandler(appAccount.AddFunds{}, &addFundsHandler); err != nil {
		return err
	}

	deductFundsHandler := appAccount.DeductFundsHandler{
		Repo:     accountRepo,
		EventBus: eventBus,
	}
	if err := cmdDispatcher.RegisterHandler(appAccount.DeductFunds{}, &deductFundsHandler); err != nil {
		return err
	}

	creatOfferHandler := appOffer.CreateOfferHandler{
		Repo:     offerRepo,
		EventBus: eventBus,
	}
	if err := cmdDispatcher.RegisterHandler(appOffer.CreateOffer{}, &creatOfferHandler); err != nil {
		return err
	}

	placeBetHandler := appOffer.PlaceBetHandler{
		Repo:     offerRepo,
		EventBus: eventBus,
	}

	if err := cmdDispatcher.RegisterHandler(appOffer.PlaceBet{}, &placeBetHandler); err != nil {
		return err
	}

	return nil
}
