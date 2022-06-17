package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

type Person struct {
	name string
	age  int
}

var p Person
var mu sync.Mutex
var vp atomic.Value

func TestNormal(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		name, age := fmt.Sprintf("nobody:%v", i), i
		go func() {
			defer wg.Done()
			p.name = name
			time.Sleep(time.Millisecond)
			p.age = age
		}()
	}
	wg.Wait()
	fmt.Printf("%+v\n", p)
}

func TestMutex(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		name, age := fmt.Sprintf("nobody:%v", i), i
		go func() {
			mu.Lock()
			defer mu.Unlock()
			defer wg.Done()
			p.name = name
			time.Sleep(time.Millisecond)
			p.age = age
		}()
	}
	wg.Wait()
	fmt.Printf("%+v\n", p)
}

func TestAtomicValue(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		name, age := fmt.Sprintf("nobody:%v", i), i
		go func() {
			defer wg.Done()
			tmp := &Person{}
			tmp.name = name
			time.Sleep(time.Millisecond)
			tmp.age = age
			vp.Store(tmp)
		}()
	}
	wg.Wait()
	fmt.Printf("%+v\n", vp.Load().(*Person))
}
