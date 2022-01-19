package main

import (
	"fmt"
	calc "modularity-demo/calculator"
	"modularity-demo/calculator/utils"
)

func main() {
	result := calc.Add(10, 10)
	fmt.Println(result)
	result = calc.Subtract(10, 10)
	fmt.Println(result)
	fmt.Println(calc.GetOpCounter())
	fmt.Println(calc.OpCounter)
	fmt.Println(utils.IsPrime(17))
}
