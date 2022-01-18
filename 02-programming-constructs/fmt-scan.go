package main

import "fmt"

func main() {
	var name string
	fmt.Printf("Enter name:")
	fmt.Scanf("%s\n", &name)
	fmt.Printf("Hello %s\n", name)

	var no int
	fmt.Printf("Enter no:")
	fmt.Scanf("%d\n", &no)
	fmt.Printf("Age:%d\n", no)

	var n1, n2 int
	fmt.Printf("Enter nos:")
	fmt.Scanf("%d %d\n", &n1, &n2)
	fmt.Printf("nos:%d and %d\n", n1, n2)

}
