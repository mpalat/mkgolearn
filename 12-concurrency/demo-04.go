package main

import (
	"fmt"
	"sync"
)

var opCount int
var mutex sync.Mutex

func main() {
	wg := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go add(i, i+1, wg)
	}
	wg.Wait()
	fmt.Println(opCount)
}

func add(x, y int, wg *sync.WaitGroup) {
	mutex.Lock()
	{
		opCount++
	}
	mutex.Unlock()
	// fmt.Println(x + y)
	wg.Done()
}
