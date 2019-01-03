package mango

import (
	"context"
	"reflect"
	"strings"

	"github.com/entropyx/tools/reflectutils"
	"github.com/entropyx/tools/strutils"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

var clientKey = &mongo.Client{}

func SetContext(c context.Context, model interface{}) error {
	doc := getDocument(model)
	doc.Context = c
	return nil
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

func structToBsonDoc(v reflect.Value) bson.D {
	doc := bson.D{}
	n := v.NumField()
	t := v.Type()
	for i := 0; i < n; i++ {
		var newFieldValue interface{}
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
		switch deepValue.Kind() {
		case reflect.Invalid:
			newFieldValue = bson.D{}
		case reflect.Struct:
			newFieldValue = structToBsonDoc(deepValue)
		default:
			newFieldValue = deepValue.Interface()
		}
		element := bson.E{newFieldName, newFieldValue}
		doc = append(doc, element)
	}
	return doc
}

func toBsonDoc(model interface{}) bson.D {
	v := reflectutils.DeepValue(reflect.ValueOf(model))
	return valueToBsonDoc(v)
}

func valueToBsonDoc(v reflect.Value) bson.D {
	doc := bson.D{}
	k := v.Kind()
	switch k {
	case reflect.Struct:
		doc = structToBsonDoc(v)
	}
	return doc
}
