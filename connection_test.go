package mango

import (
	"context"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRegister(t *testing.T) {
	Convey("Given a connection", t, func() {
		conn := &Connection{}
		Convey("When a document is registered", func() {
			s := &TestStruct{A: "1", a: "2"}
			c := context.Background()
			conn.Register(c, s)

			Convey("The document should contain the connection", func() {
				So(s.Connection(), ShouldNotBeNil)
			})
		})
	})
}
