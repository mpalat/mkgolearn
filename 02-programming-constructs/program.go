package main

import "fmt"

func main() {

	// if
	if no := 21; no%2 == 0 { // these variable have scope of the if statement blocks
		fmt.Printf("%d", no)
	} else {
		fmt.Printf("else:%d", no)
	}
	fmt.Println("")

	// for - variation 1
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Printf("sum is : %v\n", sum)

	// for - 2 - use like while
	nSum := 1
	for nSum < 100 {
		nSum += nSum
	}
	fmt.Printf("nSum is : %v\n", nSum)

	//variation 3 - infinite loop
	x := 1
	for {
		if x > 100 {
			break
		}
		x += x
	}
	fmt.Printf("x is : %v\n", x)

}
