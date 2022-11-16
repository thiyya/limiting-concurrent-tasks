package main

import (
	"flag"
	"md5-response/processor"
)

func main() {
	maxParallelCounter := flag.Int("parallel", 10, "max parallel count")
	flag.Parse()
	urls := flag.Args()
	if *maxParallelCounter > 10 {
		*maxParallelCounter = 10
	}

	processor.Processor(urls, *maxParallelCounter)
}
