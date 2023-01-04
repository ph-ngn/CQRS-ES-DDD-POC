package di

import (
	"github.com/andyj29/wannabet/internal/api/controller"
	"github.com/andyj29/wannabet/internal/command/dispatcher"
	"github.com/andyj29/wannabet/internal/command/handler"
	"github.com/andyj29/wannabet/internal/domain/account"
	"github.com/andyj29/wannabet/internal/domain/offer"
	"github.com/andyj29/wannabet/internal/eventbus"
	"github.com/andyj29/wannabet/internal/log"
	"github.com/andyj29/wannabet/internal/messaging"
	"github.com/andyj29/wannabet/internal/oauth"
	"github.com/andyj29/wannabet/internal/repository"
	"github.com/andyj29/wannabet/internal/storage"
	"github.com/segmentio/kafka-go"
	"go.uber.org/dig"
	"os"
)

func BuildContainer() *dig.Container {
	container := dig.New()
	if err := ProvideRepository(container); err != nil {
		log.GetLogger().Errorf("Failed to provide repository to di container %v", err)
	}

	if err := ProvideService(container); err != nil {
		log.GetLogger().Errorf("Failed to provide service to di container %v", err)
	}
	
	if err := ProvideAPI(container); err != nil {
		log.GetLogger().Errorf("Failed to provide api to di container %v", err)
	}

	return container
}

func ProvideAPI(container *dig.Container) error {
	_ = container.Provide(dispatcher.NewInMemoryDispatcher)
	_ = container.Provide(controller.NewAccountController)
	_ = container.Provide(func(d dispatcher.Interface) *controller.OfferController {
		return controller.NewOfferController(d, oauth.ParseSub)
	})

	return nil
}

func ProvideRepository(container *dig.Container) error {
	_ = container.Provide(func() *storage.EventStore {
		return storage.NewEventStore(os.Getenv("EVENT_STORE_ADDR"))
	})
	_ = container.Provide(repository.New[*account.Account])
	_ = container.Provide(repository.New[*offer.Offer])

	return nil
}

func ProvideService(container *dig.Container) error {
	_ = container.Provide(func() *kafka.Writer {
		return messaging.NewProducer(os.Getenv("KAFKA_ADDR"))
	})
	_ = container.Provide(eventbus.NewAsyncEventBus)
	_ = container.Provide(func(r repository.Interface[*account.Account], b eventbus.Interface) *handler.AddFundsHandler {
		return &handler.AddFundsHandler{
			Repo:     r,
			EventBus: b,
		}
	})
	_ = container.Provide(func(r repository.Interface[*account.Account], b eventbus.Interface) *handler.RegisterAccountHandler {
		return &handler.RegisterAccountHandler{
			Repo:     r,
			EventBus: b,
		}
	})
	_ = container.Provide(func(r repository.Interface[*account.Account], b eventbus.Interface) *handler.DeductFundsHandler {
		return &handler.DeductFundsHandler{
			Repo:     r,
			EventBus: b,
		}
	})
	_ = container.Provide(func(r repository.Interface[*offer.Offer], b eventbus.Interface) *handler.CreateOfferHandler {
		return &handler.CreateOfferHandler{
			Repo:     r,
			EventBus: b,
		}
	})
	_ = container.Provide(func(r repository.Interface[*offer.Offer], b eventbus.Interface) *handler.PlaceBetHandler {
		return &handler.PlaceBetHandler{
			Repo:     r,
			EventBus: b,
		}
	})
	return nil
}
