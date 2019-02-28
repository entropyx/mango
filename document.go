package mango

import (
	"context"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type D bson.D
type M bson.M

type Document struct {
	ID        primitive.ObjectID `bson:"_id"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Context   context.Context
}

func (m *Document) SetContext(c context.Context) {
	m.Context = c
}

func (m *Document) GetContext() context.Context {
	return m.Context
}

func (d *Document) Connection() *Connection {
	v := d.Context.Value("connection")
	return v.(*Connection)
}

func (d *Document) collection(model interface{}) *mongo.Collection {
	return d.Connection().collection(model)
}

func getDocument(iface interface{}) *Document {
	v := reflect.ValueOf(iface)
	if k := v.Kind(); k != reflect.Ptr {
		panic("should be a pointer")
	}
	el := v.Elem()
	docField := el.FieldByName("Document")

	return docField.Addr().Interface().(*Document)
}
