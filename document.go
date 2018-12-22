package mango

import (
	"context"

	"gopkg.in/mgo.v2/bson"
)

type D bson.D
type M bson.M

type Document struct {
	ID      bson.ObjectId
	Context context.Context
}

func (m *Document) SetContext(c context.Context) {
	m.Context = c
}

func (m *Document) GetContext() context.Context {
	return m.Context
}
