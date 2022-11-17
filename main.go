package main

import (
	"limiting-concurrent-tasks/processor"
	"limiting-concurrent-tasks/utils"
)

func main() {
	maxParallelCounter, urls := utils.HandleFlag()
	processor.Processor(urls, maxParallelCounter)
}
