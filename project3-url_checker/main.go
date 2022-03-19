package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request failed")

func main() {
	var results = make(map[string]string)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://soundcloud.com/",
	}
	for _, url := range urls {
		// fmt.Println(url)
		result := "OK"
		err := hitURL(url)
	}
}

func hitURL(url string) error {
	fmt.Println("Checking ", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		return errRequestFailed
	} else {
		fmt.Println(resp.StatusCode)
	}
	return nil
}
