package common

import "github.com/andyj29/wannabet/internal/domain/common"

type Repository[T common.AggregateRoot] interface {
	Load(string) (T, error)
	Save(T) error
}
