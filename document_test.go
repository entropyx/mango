package mango

import (
	"context"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_getDocument(t *testing.T) {
	Convey("Given a struct with a reference to a document", t, func() {
		s := &TestStruct{}

		Convey("When the document is got", func() {
			doc := getDocument(s)

			Convey("doc should be valid", func() {
				So(doc, ShouldNotBeNil)
				So(doc.Context, ShouldBeNil)
			})

			Convey("And the context is modified", func() {
				doc.Context = context.Background()

				Convey("The new context should be valid", func() {
					newDoc := getDocument(s)
					So(newDoc.Context, ShouldNotBeNil)
				})
			})
		})
	})
}
