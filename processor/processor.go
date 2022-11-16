package processor

import (
	"fmt"
	"log"
	"md5-response/utils"
	"sync"
)

// Processor Processes urls by limiting the number of concurrent tasks running
func Processor(urls []string, maxParallelCounter int) {
	var waitGroup sync.WaitGroup

	limiter := make(chan bool, maxParallelCounter)
	for _, url := range urls {
		waitGroup.Add(1)
		go process(url, limiter, &waitGroup)
	}

	waitGroup.Wait()
}

func process(url string, limiter chan bool, waitGroup *sync.WaitGroup) {
	defer func() {
		<-limiter
		waitGroup.Done()
	}()
	limiter <- true
	fixedUrl := utils.FixUrl(url)
	b, err := utils.GetResponse(fixedUrl)
	if err != nil {
		log.Printf("Response not received from %s --> err : %v", fixedUrl, err)
		return
	}
	result, err := utils.ByteToMd5(b)
	if err != nil {
		log.Printf("Error occured while converting byte to Md5 --> err : %v", err)
		return
	}
	fmt.Println(fmt.Sprintf("%s - %s", fixedUrl, result))
}
