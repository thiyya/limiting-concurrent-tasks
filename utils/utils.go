package utils

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	prefix_s = "https://"
	prefix   = "http://"
)

// ByteToMd5 Converts byte slice to md5 hash
func ByteToMd5(b []byte) (string, error) {
	return fmt.Sprintf("%x", md5.Sum(b)), nil
}

// GetResponse Returns response of the url as byte slice
func GetResponse(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// FixUrl Adds prefix to prevent unsupported protocol scheme err
func FixUrl(url string) string {
	if !strings.HasPrefix(url, prefix) && !strings.HasPrefix(url, prefix_s) {
		return fmt.Sprintf(prefix_s + url)
	}

	return url
}

// HandleFlag Gets max parallel counter and url slice from args
// Default max parallel counter is 10
func HandleFlag() (int, []string) {
	maxParallelCounter := flag.Int("parallel", 10, "max parallel count")
	flag.Parse()
	urls := flag.Args()
	if *maxParallelCounter > 10 {
		*maxParallelCounter = 10
	}
	return *maxParallelCounter, urls
}
