package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go f1()
	f2()
	wg.Wait()
	// var input string
	// fmt.Scanln(&input) - wait for the input

	// time.Sleep(1 * time.Millisecond) - sleeping

}

func f1() {
	fmt.Println("f1")
	wg.Done()
}
func f2() {
	fmt.Println("f2")
}
