package utils

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestUtils(t *testing.T) {
	Convey("GetResponse", t, func(c C) {
		b, err := GetResponse("https://google.com")
		c.So(err, ShouldBeNil)
		c.So(b, ShouldNotBeNil)
		c.So(string(b), ShouldContainSubstring, "<title>Google</title>")
	})

	Convey("ByteToMd5", t, func(c C) {
		str := "byteSlice"
		s, err := ByteToMd5([]byte(str))
		c.So(err, ShouldBeNil)
		c.So(s, ShouldNotBeNil)
		c.So(len(s), ShouldBeGreaterThan, len(str))
	})

	Convey("FixUrl", t, func(c C) {
		c.So(FixUrl("https://google.com"), ShouldEqual, "https://google.com")
		c.So(FixUrl("http://google.com"), ShouldEqual, "http://google.com")
		c.So(FixUrl("google.com"), ShouldEqual, "https://google.com")
	})
}
