package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)
	evenNoCh := genEvenNos(done)
	go func() {
		for evenNo := range evenNoCh {
			fmt.Println("evenNo : ", evenNo)
		}
	}()
}

func genEvenNos(doneChannel chan bool) <-chan int {
	resultCh := make(chan int)
	go func(done chan bool) {
		var no int
		for {
			select {
			case <-done:
				fmt.Println("done")
				close(resultCh)

			case resultCh <- no:
				time.Sleep(500 * time.Millisecond)
				no += 2

			}
		}
	}(doneChannel)
	return resultCh
}
