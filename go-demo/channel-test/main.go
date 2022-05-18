package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

type hchan struct {
	qcount   uint           // total data in the queue
	dataqsiz uint           // size of the circular queue
	buf      unsafe.Pointer // points to an array of dataqsiz elements
	elemsize uint16
	closed   uint32
	elemtype unsafe.Pointer // element type
	sendx    uint           // send index
	recvx    uint           // receive index
	recvq    [2]uintptr     // list of recv waiters
	sendq    [2]uintptr     // list of send waiters
	lock     [3]uintptr
}

type emptyInterface struct {
	_type unsafe.Pointer
	value unsafe.Pointer
}

func getPointer(m interface{}) unsafe.Pointer {
	ei := (*emptyInterface)(unsafe.Pointer(&m))
	return ei.value
}

func printStruct() {
	h := make(chan int, 1)
	hm := (*hchan)(getPointer(h))
	fmt.Printf("%+v\n", hm)
}

func writeToNilChan() {
	var h chan int
	go func() {
		h = make(chan int)
		<-h
	}()
	h <- 1
}

func readFromCloseChan() {
	h := make(chan int, 1)
	h <- 2
	close(h)
	a := <-h
	print(a)
}

func nSenderNReceiver() {
	const MAX = 1000
	numCount := int64(0)
	oneTime := sync.Once{}

	dataCh := make(chan int, 100)
	stopCh := make(chan struct{})

	// senders
	for i := 0; i < 10; i++ {
		go func(a int) {
			for {
				select {
				case <-stopCh:
					return
				case dataCh <- a:
				}
			}
		}(i)
	}

	// the receiver
	for i := 0; i < 10; i++ {
		go func() {
			for {
				select {
				case <-stopCh:
					return
				case v := <-dataCh:
					if atomic.AddInt64(&numCount, 1) >= MAX {
						oneTime.Do(func() {
							close(stopCh)
						})
					}
					println(v)
				}
			}
		}()
	}

	time.Sleep(time.Hour)
}

func main() {
	nSenderNReceiver()
}
