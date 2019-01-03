package mango

import (
	"context"
	"os"
	"testing"

	"github.com/mongodb/mongo-go-driver/bson"
	. "github.com/smartystreets/goconvey/convey"
)

type TestStruct struct {
	Document
	A     string
	a     string
	Child *TestStruct
}

func Test_getCollection(t *testing.T) {
	Convey("Given a document", t, func() {
		s := &TestStruct{}

		Convey("When the collection name is got", func() {
			name := getCollection(s)

			Convey("The collection name should not be empty", func() {
				So(name, ShouldEqual, "test_struct")
			})
		})
	})
}

func Test_toBsonDoc(t *testing.T) {
	Convey("Given a model", t, func() {
		s := &TestStruct{A: "1", a: "2"}
		Convey("When it is converted to a document", func() {
			doc := toBsonDoc(s)

			Convey("It should be valid", func() {
				So(doc, ShouldResemble, bson.D{{"a", "1"}, {"child", bson.D{}}})
			})
		})
	})
}

func TestUpsertOne(t *testing.T) {
	Convey("Given a model", t, WithConnection(func(conn *Connection) {
		s := &TestStruct{A: "1", Child: &TestStruct{A: "2"}}
		conn.Register(context.Background(), s)
		Convey("When it is upserted", func() {
			err := UpdateOne(bson.D{{"name", "Ash"}}, Set(s))

			Convey("err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("The id should be valid", func() {

			})
		})
	}))
}

func WithConnection(f func(conn *Connection)) func() {
	return func() {
		config := &Config{
			Address:  os.Getenv("MONGO_HOST"),
			Port:     27017,
			Database: "mango_test",
		}
		conn, err := Connect(config)
		if err != nil {
			panic(err)
		}
		Reset(func() {

		})
		f(conn)
	}
}
