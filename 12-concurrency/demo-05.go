package main

import (
	"fmt"
	"sync"
)

type opLock struct {
	count int
	sync.Mutex
}

var opLocker opLock

func main() {
	wg := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go add(i, i+1, wg)
	}
	wg.Wait()
	fmt.Println(opLocker.count)
}

func add(x, y int, wg *sync.WaitGroup) {
	opLocker.Lock()
	{
		opLocker.count++
	}
	opLocker.Unlock()
	// fmt.Println(x + y)
	wg.Done()
}
