package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	prefix = "https://"
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
	if !strings.HasPrefix(url, prefix) {
		return fmt.Sprintf(prefix + url)
	}

	return url
}
