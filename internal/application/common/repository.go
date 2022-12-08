package common

import "github.com/andyj29/wannabet/internal/domain/common"

type Repository interface {
	Load(string) common.AggregateRoot
	Save(common.AggregateRoot) error
}
