package common

import "reflect"

type Query interface {
	GetQueryType() string
}

type QueryBase struct {
}

func (q *QueryBase) GetQueryType() string {
	return reflect.TypeOf(q).Elem().Name()
}
