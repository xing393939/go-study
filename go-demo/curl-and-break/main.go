package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func curl() {
	client := http.Client{Timeout: time.Millisecond * 10}
	response, err := client.Get("http://192.168.2.119:8005/read5/select")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(response.Status)
	}
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				curl()
			}
		}()
	}
	wg.Wait()
}
