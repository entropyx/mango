package mango

import (
	"context"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/mgo.v2/bson"
)

type testStruct struct {
	Document
	A     string
	a     string
	Child *testStruct
}

func Test_toBsonDoc(t *testing.T) {
	Convey("Given a model", t, func() {
		s := &testStruct{A: "1", a: "2"}
		Convey("When it is converted to a document", func() {
			doc := toBsonDoc(s)

			Convey("It should be valid", func() {
				So(doc, ShouldResemble, bson.D{{"a", "1"}, {"child", bson.D{}}})
			})
		})
	})
}

func TestSetContext(t *testing.T) {
	Convey("Given a context with a value", t, func() {
		c := context.Background()
		c = context.WithValue(c, "test", "test")
		Convey("When the context is set to a document", func() {
			s := &testStruct{A: "1", a: "2"}
			err := SetContext(c, s)

			Convey("The document should contain a new context")
		})
	})
}

func TestUpsertOne(t *testing.T) {
	Convey("Given a model", t, func() {
		s := &testStruct{A: "1", Child: &testStruct{A: "2"}}

		Convey("When it is upserted", func() {
			UpsertOne(D{{"_id", bson.NewObjectId()}}, s)

			Convey("The id should be valid", func() {

			})
		})
	})
}
