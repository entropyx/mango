package mango

import (
	"context"
	"reflect"
	"strings"

	"github.com/entropyx/mango/options"
	"github.com/entropyx/tools/reflectutils"
	"github.com/entropyx/tools/strutils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	opts "go.mongodb.org/mongo-driver/mongo/options"
)

var clientKey = &mongo.Client{}
var keyConnection *Connection

func SetContext(c context.Context, model interface{}) error {
	doc := getDocument(model)
	doc.Context = c
	return nil
}

func FindOne(filter interface{}, value interface{}, ops ...*options.FindOne) error {
	doc := getDocument(value)
	collection := doc.collection(value)
	// TODO: set to mongo options
	result := collection.FindOne(doc.Context, filter, &opts.FindOneOptions{})
	return result.Decode(value)
}

func InsertOne(value interface{}, ops ...*options.InsertOne) error {
	doc := getDocument(value)
	collection := doc.collection(value)
	// TODO: set to mongo options
	result, err := collection.InsertOne(doc.Context, toBsonDoc(value))
	if err != nil {
		return err
	}
	id := result.InsertedID.(primitive.ObjectID)
	doc.ID = id
	return nil
}

func UpdateOne(filter interface{}, operator *Operator, ops ...*options.Update) error {
	var updateOptions []*opts.UpdateOptions
	doc := getDocument(operator.Value)
	collection := doc.collection(operator.Value)
	for _, op := range ops {
		updateOptions = append(updateOptions, &opts.UpdateOptions{Upsert: &op.Upsert})
	}
	_, err := collection.UpdateOne(doc.Context, filter, operator.apply(), updateOptions...)
	return err
}

func getContextFromModel(model interface{}) context.Context {
	doc := getDocument(model)
	return doc.Context
}

func getCollection(i interface{}) string {
	t := reflect.TypeOf(i)
	split := strings.Split(t.String(), ".")
	name := split[len(split)-1]
	snakedName := strutils.ToSnakeCase(name)
	return snakedName
}

func arrayToBsonA(v reflect.Value) bson.A {
	l := v.Len()
	a := bson.A{}
	for i := 0; i < l; i++ {
		itemValue := v.Index(i)
		deepValue := reflectutils.DeepValue(itemValue)
		item := valueToBson(deepValue)
		a = append(a, item)
	}
	return a
}

func structToBsonDoc(v reflect.Value) bson.D {
	doc := bson.D{}
	n := v.NumField()
	t := v.Type()
	for i := 0; i < n; i++ {
		field := t.Field(i)
		fieldName := field.Name
		newFieldName := strutils.ToSnakeCase(fieldName)
		fieldValue := v.Field(i)
		if !fieldValue.CanInterface() {
			continue
		}
		// TODO: this looks ugly
		if ft := fieldValue.Type(); ft.String() == "mango.Document" {
			// TODO: document fields
			continue
		}
		deepValue := reflectutils.DeepValue(fieldValue)
		newFieldValue := valueToBson(deepValue)
		element := bson.E{newFieldName, newFieldValue}
		doc = append(doc, element)
	}
	return doc
}

func toBsonDoc(model interface{}) bson.D {
	v := reflectutils.DeepValue(reflect.ValueOf(model))
	return valueToBson(v).(bson.D)
}

func valueToBson(v reflect.Value) interface{} {
	k := v.Kind()
	switch k {
	case reflect.Invalid:
		return bson.D{}
	case reflect.Struct:
		return structToBsonDoc(v)
	case reflect.Array, reflect.Slice:
		return arrayToBsonA(v)
	default:
		return v.Interface()
	}
	return bson.D{}
}
