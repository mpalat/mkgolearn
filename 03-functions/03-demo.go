// variadic functions

package main

import "fmt"

func main() {

	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sumDum("Dummy", 1, 2, 3, 4, 5))
}

func sum(nos ...int) int {
	result := 0
	for i := 0; i < len(nos); i++ {
		result += nos[i]
	}
	return result
}

func sumDum(s string, nos ...int) int {
	result := 0
	for i := 0; i < len(nos); i++ {
		result += nos[i]
	}
	fmt.Println(s)
	return result
}
