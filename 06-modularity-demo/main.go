package main

import (
	"fmt"
	"modularity-demo/calculator"
	"modularity-demo/calculator/utils"
)

func main() {
	result := calculator.Add(10, 10)
	fmt.Println(result)
	fmt.Println(utils.IsPrime(result))
}
