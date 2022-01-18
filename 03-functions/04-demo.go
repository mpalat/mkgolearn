//Higher order functions

package main

import "fmt"

func main() {

	fn := func() {
		fmt.Println("fn invoked")
	}
	fn()

	var ffn = func() {
		fmt.Println("with var")
	}
	ffn()

	var vfn func()
	vfn = func() {
		fmt.Println("vfn")
	}
	vfn()

	var add func(int, int) int
	add = func(x, y int) int {
		return x + y
	}
	fmt.Println(add(100, 10))
}

func performOperation(oper func(int, int) int, x, y int) {
	// return oper(x, y)
}
