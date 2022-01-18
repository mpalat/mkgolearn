package main

import "fmt"

func main() {
	if no := 21; no%2 == 0 { // these variable have scope of the if statement blocks
		fmt.Printf("%d", no)
	} else {
		fmt.Printf("else:%d", no)
	}
	fmt.Println("")
}
