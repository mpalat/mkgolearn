package main

import "fmt"

func main() {
	var no int = 10
	// var addrNo = &no
	var addrNo *int = &no
	// fmt.Println(no, addrNo)
	fmt.Println(*addrNo, addrNo)
	fmt.Println("after incrementing")
	increment(addrNo)
	fmt.Printf("no:%d, addrNo:%p\n", no, addrNo)

	var n1, n2 = 10, 20
	fmt.Printf("Before swapping - n1 = %d, n2 = %d\n", n1, n2)
	swap(&n1, &n2)
	fmt.Printf("After swapping - n1 = %d, n2 = %d\n", n1, n2)
}

func increment(addrNo *int) {
	*addrNo++
}
func swap(px, py *int) {
	*px, *py = *py, *px
}
