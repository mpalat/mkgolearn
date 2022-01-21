package main

import (
	"fmt"
	"time"
)

func main() {
	evenNoCh := genEvenNos()
	for i := 0; i < 20; i++ {
		evenNo := <-evenNoCh
		fmt.Println("evenNo : ", evenNo)
	}
}

func genEvenNos() <-chan int {
	resultCh := make(chan int)
	go func() {
		var no int
		for {
			time.Sleep(500 * time.Millisecond)
			no += 2
			fmt.Println("Generated : ", no)
			resultCh <- no
		}
	}()
	return resultCh
}
