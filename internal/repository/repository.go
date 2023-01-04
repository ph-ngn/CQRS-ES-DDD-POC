package repository

import "github.com/andyj29/wannabet/internal/domain/common"

type Interface[T common.AggregateRoot] interface {
	Load(string) (T, error)
	Save(T) error
}
