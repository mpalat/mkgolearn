package main

import "fmt"

func main() {

	fmt.Println("Primes")
	for i := 3; i < 100; i++ {
		var maxDivider = i / 3

		var prime = true
		for j := 2; j <= maxDivider; j++ {
			if i%j == 0 {
				prime = false
				break
			}
		}
		if prime == true {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println("")

}
