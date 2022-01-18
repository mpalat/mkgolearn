package main

import "fmt"

func main() {
	fmt.Println(isPrime(13))
	fmt.Println(add(12, 43))
	fmt.Println(divide(10, 3))
	quotient, remainder := divide(10, 3)
	fmt.Printf("quotient : %d, remainder:%d\n", quotient, remainder)
	quotient1, _ := divide(10, 3)
	fmt.Printf("quotient : %d\n", quotient1)

	fmt.Println(divide2(10, 3))

}

func isPrime(no int) bool {
	var prime = true
	var maxDivider int = no / 2
	for j := 2; j <= maxDivider; j++ {
		if no%j == 0 {
			prime = false
			break
		}
	}
	return prime
}

// named return value
func add(x, y int) (result int) {
	result = x + y
	return
}

// more than one result
func divide(x, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}
func divide2(x, y int) (quotient, remainder int) {
	quotient = x / y
	remainder = x % y
	return
}
