package main

import (
	"fmt"
	calc "modularity-demo/calculator"

	// . "modularity-demo/calculator" - . to import globally - not advised
	"modularity-demo/calculator/utils"

	_ "modularity-demo/dummy" // _ will make sure init fn will be called even though there is no use

	"github.com/fatih/color"
)

func main() {
	fmt.Println("hello")
	result := calc.Add(10, 10)
	fmt.Println(result)
	result = calc.Subtract(10, 10)
	fmt.Println(result)
	fmt.Println(calc.GetOpCounter())
	fmt.Println(calc.OpCounter)
	fmt.Println(utils.IsPrime(17))
	color.Blue("is 97 a prime number:%t\n", utils.IsPrime(97))
}
