package main

import (
	"fmt"
)

// share by communicating

func main() {
	// var ch chan int = make(chan int)
	ch := make(chan int)

	go func() {
		ch <- 100
	}()
	result := <-ch
	go chWrite(ch)
	result = <-ch
	fmt.Println("after wait:", result)
}

func chWrite(ch chan int) {
	ch <- 100

}
