package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func httpDo(ctx context.Context, req *http.Request, f func(*http.Response, error) error) error {
	// Run the HTTP request in a goroutine and pass the response to f.
	c := make(chan error, 1)
	req = req.WithContext(ctx)
	go func() {
		c <- f(http.DefaultClient.Do(req))
	}()
	select {
	case <-ctx.Done():
		<-c
		return ctx.Err()
	case err := <-c:
		return err
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	req, err := http.NewRequest("GET", "https://httpbin.org/delay/2", nil)
	if err != nil {
		return
	}
	err = httpDo(ctx, req, func(resp *http.Response, e error) error {
		if err != nil {
			return err
		}

		if resp != nil {
			bytes, _ := ioutil.ReadAll(resp.Body)
			println(string(bytes))
		}
		return nil
	})
	fmt.Println(err)
}
