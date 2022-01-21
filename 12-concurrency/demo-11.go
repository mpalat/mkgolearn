package main

import (
	"fmt"
)

/*
	Assignment:
		Print "Hello" & "World" simultaneously
		IMPORTANT : Do NOT move the loop outside the "print" function
*/

func main() {
	ch1 := make(chan bool)
	ch2 := make(chan bool, 1)
	go print("Hello", ch1, ch2)
	go print("World", ch2, ch1)
	ch1 <- true
	fmt.Println("End of main!")
}

func print(s string, ch1, ch2 chan bool) {
	for i := 0; i < 5; i++ {
		<-ch1
		fmt.Println(s)
		ch2 <- true
		// time.Sleep(500 * time.Millisecond)
	}
}
