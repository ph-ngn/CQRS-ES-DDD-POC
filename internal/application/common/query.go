package common

import "reflect"

type QueryModel interface {
	GetQueryModelType() string
}

type QueryModelBase struct {
}

func (q *QueryModelBase) GetQueryModelType() string {
	return reflect.TypeOf(q).Elem().Name()
}
