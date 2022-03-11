package main

import "time"

func bigCache() {
	num := 1024 * 1024 * 1024 * 2
	arr := make([]byte, num, num)
	i := 1
	for {
		i++
		arr[i%num] = byte(1)
	}
}

func timeLeak() {
	for {
		select {
		case <-time.NewTicker(time.Nanosecond).C:
			_ = 1
		case <-time.NewTicker(time.Nanosecond * 2).C:
			_ = 2
		}
	}
}

func main() {
	go bigCache()
	//go timeLeak()
	for {
	}
}
