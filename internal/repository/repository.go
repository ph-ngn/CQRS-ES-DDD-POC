package repository

import (
	"github.com/andyj29/wannabet/internal/domain"
	"reflect"

	"github.com/andyj29/wannabet/internal/log"
	"github.com/andyj29/wannabet/internal/storage"
)

type Interface[T domain.AggregateRoot] interface {
	Load(string) (T, error)
	Save(T) error
}

type Repository[T domain.AggregateRoot] struct {
	es *storage.EventStore
}

func New[T domain.AggregateRoot](eventStore *storage.EventStore) *Repository[T] {
	return &Repository[T]{
		es: eventStore,
	}
}

func (r *Repository[T]) Load(aggregateID string) (T, error) {
	var deserializedAggregate T
	initNilAggregatePtr(&deserializedAggregate)

	stream, errStream := r.es.ReadAll(aggregateID)
	for {
		select {
		case event := <-stream:
			deserializedAggregate.When(event, false)
		case err := <-errStream:
			if err != nil {
				var nilAggregate T
				return nilAggregate, err
			}
			return deserializedAggregate, nil
		}
	}
}

func (r *Repository[T]) Save(aggregate T) error {
	changes := aggregate.GetChanges()
	for _, change := range changes {
		if err := r.es.Append(change, nil); err != nil {
			log.GetLogger().Errorf(err.Error())
			return err
		}
	}
	return nil
}

func initNilAggregatePtr(dp interface{}) {
	target := reflect.ValueOf(dp).Elem()
	if reflect.Indirect(target).IsValid() {
		return
	}

	aggregateType := target.Type().Elem()
	target.Set(reflect.New(aggregateType))
}
