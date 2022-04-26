package main

import (
	"context"
	"errors"
	"golang.org/x/sync/errgroup"
	"time"
)

func breakWhenHasError() {
	wg, ctx := errgroup.WithContext(context.Background())
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
	}
	for sec, url := range urls {
		url := url
		sec := sec
		wg.Go(func() error {
			i := 0
			for {
				select {
				case <-ctx.Done():
					return nil
				default:
					time.Sleep(time.Second)
					println(time.Now().Format("15:04:05"), url)
					if i > sec {
						return errors.New(url)
					}
					i++
				}
			}
		})
	}
	err := wg.Wait()
	if err == nil {
		println("Successfully fetched all URLs.")
	} else {
		println("breakWhenHasError:", err.Error())
	}
}

func noBreakWhenHasError() {
	var wg errgroup.Group
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
	}
	for sec, url := range urls {
		url := url
		sec := sec
		wg.Go(func() error {
			time.Sleep(time.Second * time.Duration(sec))
			println(time.Now().Format("15:04:05"), url)
			return errors.New(url)
		})
	}
	err := wg.Wait()
	if err == nil {
		println("Successfully fetched all URLs.")
	} else {
		println("noBreakWhenHasError:", err.Error())
	}
}

func main() {
	breakWhenHasError()
	noBreakWhenHasError()
}
