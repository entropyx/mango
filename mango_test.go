package mango

import (
	"context"
	"os"
	"testing"

	"github.com/entropyx/mango/options"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	. "github.com/smartystreets/goconvey/convey"
)

type TestStruct struct {
	Document `bson:",inline"`
	A        string
	a        string
	Child    *TestStruct
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

func TestInsertOne(t *testing.T) {
	Convey("Given a document", t, WithConnection(func(conn *Connection) {
		s := &TestStruct{A: "1", Child: &TestStruct{A: "2"}}
		conn.Register(context.Background(), s)

		Convey("When it is inserted", func() {
			err := InsertOne(s)

			Convey("err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("The id should NOT be zero", func() {
				So(s.ID.IsZero(), ShouldBeFalse)
			})
		})
	}))
}

func TestUpsertOne(t *testing.T) {
	Convey("Given a document", t, WithConnection(func(conn *Connection) {
		s := &TestStruct{A: "1", Child: &TestStruct{A: "2"}}
		conn.Register(context.Background(), s)

		Convey("When it is upserted and it is NOT found", func() {
			id := primitive.NewObjectID()
			err := UpdateOne(bson.D{{"_id", id}}, Set(s), &options.Update{Upsert: true})

			Convey("err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("The document should exist now", func() {
				fs := &TestStruct{}
				conn.Register(context.Background(), fs)
				FindOne(bson.D{{"_id", id}}, fs)
				So(fs.ID.IsZero(), ShouldBeFalse)
				So(fs.ID.Hex(), ShouldEqual, id.Hex())
			})
		})

		Convey("When it is upserted and it is found", func() {
			s := &TestStruct{A: "1", Child: &TestStruct{A: "2"}}
			conn.Register(context.Background(), s)

			InsertOne(s, &options.InsertOne{})

			Convey("The document should exist now", func() {
				fs := &TestStruct{}
				conn.Register(context.Background(), fs)
				FindOne(bson.D{{"_id", s.ID}}, fs)
				So(fs.ID.IsZero(), ShouldBeFalse)
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
