package calculator

import "fmt"

var OpCounter int = 0

func init() {
	fmt.Println("calc inited")
}
func GetOpCounter() int {
	return OpCounter
}
