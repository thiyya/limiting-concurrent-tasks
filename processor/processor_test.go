package processor

import (
	. "github.com/smartystreets/goconvey/convey"
	"sync"
	"testing"
)

var waitGroup sync.WaitGroup

func TestProcess(t *testing.T) {
	Convey("Process ", t, func(c C) {
		limiter := make(chan bool, 10)
		waitGroup.Add(1)
		err := process("https://google.com", limiter, &waitGroup)
		c.So(err, ShouldBeNil)
	})
}

func TestProcessor(t *testing.T) {
	var urls = []string{
		"https://google.com",
		"https://github.com",
		"https://stackoverflow.com",
		"https://www.linkedin.com",
		"https://www.facebook.com",
		"https://twitter.com",
	}
	Convey("Processor ", t, func(c C) {
		err := Processor(urls, 10)
		c.So(err, ShouldBeNil)
	})
}
