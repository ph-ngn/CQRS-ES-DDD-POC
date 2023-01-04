package storage

import (
	"github.com/andyj29/wannabet/internal/domain"
	"github.com/andyj29/wannabet/internal/log"
	goes "github.com/jetbasrawi/go.geteventstore"
	"net/url"
	"time"
)

type EventStore struct {
	client        *goes.Client
	eventRegistry map[string]func() domain.Event
}

func NewEventStore(addr string) *EventStore {
	client, err := goes.NewClient(nil, addr)
	if err != nil {
		log.GetLogger().Fatalf("failed to establish new event store http connection")
	}

	return &EventStore{
		client: client,
	}
}

func (es *EventStore) Append(event domain.Event, meta map[string]string) error {
	streamWriter := es.client.NewStreamWriter(event.GetAggregateID())
	newEvent := goes.NewEvent(goes.NewUUID(), event.GetEventType(), event, meta)
	return streamWriter.Append(nil, newEvent)
}

func (es *EventStore) ReadAll(streamID string) (<-chan domain.Event, <-chan error) {
	stream := make(chan domain.Event)
	errStream := make(chan error)

	go func() {
		streamReader := es.client.NewStreamReader(streamID)
		for streamReader.Next() {
			if err := streamReader.Err(); err != nil {
				switch err.(type) {
				case *url.Error, *goes.ErrTemporarilyUnavailable:
					log.GetLogger().Infof("The event store server is not ready at the moment: %v. Attempt to retry after 10 seconds", err)
					<-time.After(time.Duration(10) * time.Second)

				case *goes.ErrNotFound:
					log.GetLogger().Errorf("Could not find stream with this ID: %v")
					errStream <- err
					return

				case *goes.ErrUnauthorized:
					log.GetLogger().Fatalf("Read is not authorized for this stream: %v", err)

				case *goes.ErrNoMoreEvents:
					errStream <- nil
					return
				}
			}

			var event domain.Event

			if initEvent, ok := es.eventRegistry[streamReader.EventResponse().Event.EventType]; ok {
				event = initEvent()
				if err := streamReader.Scan(&event, nil); err != nil {
					log.GetLogger().Fatalf(err.Error())
				}
				stream <- event
			}
		}
	}()

	return stream, errStream
}
