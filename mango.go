package mango

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	"github.com/entropyx/tools/reflectutils"
	"github.com/entropyx/tools/strutils"
	"github.com/mongodb/mongo-go-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

var clientKey = &mongo.Client{}

func SetContext(c context.Context, model interface{}) error {
	doc, err := getDocument(model)
	if err != nil {
		return err
	}
	doc.Context = c
	return nil
}

func getContextFromModel(model interface{}) (context.Context, error) {
	doc, err := getDocument(model)
	if err != nil {
		return nil, err
	}
	return doc.Context, nil
}

func getDocument(iface interface{}) (Document, error) {
	v := reflect.ValueOf(iface)
	if k := v.Kind(); k != reflect.Ptr {
		return Document{}, errors.New("should be a pointer")
	}
	el := v.Elem()
	docField := el.FieldByName("Document")
	doc := docField.Interface().(Document)
	return doc, nil
}

func getCollection(i interface{}) string {
	t := reflect.TypeOf(i)
	name := strutils.ToSnakeCase(t.Name())
	return name
}

func structToBsonDoc(v reflect.Value) bson.D {
	doc := bson.D{}
	n := v.NumField()
	t := v.Type()
	fmt.Println("type", t)
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
		if deepValue.Kind() == reflect.Invalid {
			newFieldValue = bson.D{}
		} else {
			newFieldValue = deepValue.Interface()
		}
		element := bson.DocElem{newFieldName, newFieldValue}
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
