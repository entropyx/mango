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

func Test_portString(t *testing.T) {
	Convey("Given a valid port number", t, func() {
		var port uint = 3000
		Convey("When the port is parsed to string", func() {
			s := portString(port)
			Convey("The string should be valid", func() {
				So(s, ShouldEqual, ":3000")
			})
		})
	})

	Convey("Given a zero port number", t, func() {
		var port uint = 0
		Convey("When the port is parsed to string", func() {
			s := portString(port)
			Convey("The string should be empty", func() {
				So(s, ShouldBeEmpty)
			})
		})
	})
}

func Test_hostString(t *testing.T) {
	Convey("Given a address and port", t, func() {
		var port uint = 3000
		host := "example.com"

		Convey("When the host string is generated", func() {
			s := hostString(host, port)
			Convey("The string should be valid", func() {
				So(s, ShouldEqual, "example.com:3000")
			})
		})
	})
}
