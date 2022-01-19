package main

import "fmt"

func main() {
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())

	inc := getIncrement()
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
}

var counter int = 0

func increment() int {
	counter++
	return counter
}

// should not be able to influence the change of counter from outside
// So we want this variable accessible only to the increment function
// Because the inner function does not go away, the outer counter is not GCed and hence alive
// ie the lifetime of the variable in closure
func getIncrement() func() int {
	var counter int
	increment := func() int {
		counter++
		return counter
	}
	return increment
}
