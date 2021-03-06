package main

import "fmt"

func main() {
	// var x int
	var x = "Hello World"
	fmt.Println(x)

	fmt.Printf("value:%v of type %T\n", x, x)

	y := 100 // only applicable inside a function
	fmt.Println(y)

	var (
		n1, n2 int = 10, 20
		result     = n1 + n2
	)
	fmt.Println(result)

	n3 := 10
	n4 := 20
	fmt.Println(n3)
	fmt.Println(n4)

	a, b, str := 10, 20, "a:%v b:%v\n"
	fmt.Printf(str, a, b)

	const z = 100

	const (
		red = iota + 5
		green
		_
		blue
	)

	fmt.Println(red, green, blue)

	// Type conversion
	var vInt int = 100
	var vFloat float32 = 200.00
	var vResult = float32(vInt) + vFloat
	fmt.Printf("vResult is %f of Type %T\n", vResult, vResult)

	var vResult1 = vInt + int(vFloat)
	fmt.Printf("vResult1 is %d of Type %T\n", vResult1, vResult1)

}
