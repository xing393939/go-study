package main

import (
	"golang.org/x/sync/errgroup"
	"net/http"
)

func ExampleGroup_justErrors() {
	var wg errgroup.Group
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
	}
	for _, url := range urls {
		url := url
		wg.Go(func() error {
			_, err := http.Get(url)
			return err
		})
	}
	// Wait for all HTTP fetches to complete.
	err := wg.Wait()
	if err == nil {
		println("Successfully fetched all URLs.")
	} else {
		println(err.Error())
	}
}

func main() {
	ExampleGroup_justErrors()
}
