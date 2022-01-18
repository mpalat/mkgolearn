// higher with functions as arguments

package main

import "fmt"

func main() {

	performOperation(add, 100, 10)
	performOperation(sub, 100, 10)
	performOperation(mul, 100, 10)
	performOperation(divide, 100, 10)
}

/*
pn invoking the operator functions print Before invocation, result and then after invocation
*/

func performOperation(oper func(int, int) int, x, y int) {

	fmt.Println("Before Operation")
	result := oper(x, y)
	fmt.Printf("result:%d\n", result)
	fmt.Println("After Operation")
}

func add(x, y int) int {
	return x + y
}
func sub(x, y int) int {
	return x - y
}
func mul(x, y int) int {
	return x * y
}
func divide(x, y int) int {
	return x / y
}
