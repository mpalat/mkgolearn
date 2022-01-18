package main

import "fmt"

func main() {

	func() {
		fmt.Println("fn invoked")
	}()

	func(x, y int) {
		fmt.Println(x + y)
	}(100, 200)

	subResult := func(x, y int) int {
		return x - y
	}(200, 10)
	fmt.Printf("subResult:%d\n", subResult)
}
