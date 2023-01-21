package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("request failed")

func main() {
	urls := []string{
		"https://airbnb.com",
		"https://www.google.com",
		"https://www.amazon.com",
		"https://www.reddit.com",
		"https://www.facebook.com",
		"https://www.youtube.com",
		"https://www.yahoo.com",
		"https://www.wikipedia.org",
		"https://www.twitter.com",
	}

	for _, url := range urls {
		hitURL(url)
	}
}

func hitURL(url string) error {
	fmt.Println(fmt.Sprint("Hitting URL: ", url))
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		return errRequestFailed
	}
	return nil
}
