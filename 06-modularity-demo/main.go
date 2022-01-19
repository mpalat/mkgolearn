package main

import (
	"fmt"
	"modularity-demo/calculator"
	"modularity-demo/calculator/utils"
)

func main() {
	result := calculator.Add(10, 10)
	fmt.Println(result)
	result = calculator.Subtract(10, 10)
	fmt.Println(result)
	fmt.Println(calculator.GetOpCounter())
	fmt.Println(calculator.OpCounter)
	fmt.Println(utils.IsPrime(17))
}
