package main

import (
	"fmt"
	"sync"
)

// communicate by sharing memory - NOT PREFERRED
var result int

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go add(10, 1, wg)
	wg.Wait()
	fmt.Println("after wait:"result++)
}

func add(x, y int, wg *sync.WaitGroup) {
	result++
	wg.Done()

}
