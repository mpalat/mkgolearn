package main

import "fmt"

func main() {

	fmt.Println("1.Add")
	fmt.Println("2.Subtract")
	fmt.Println("3.Multiply")
	fmt.Println("4.Divide")
	fmt.Println("5.Exit")
	fmt.Println("Enter a Choice:")

	var choice int
	fmt.Scanf("%d\n", &choice)

	if choice >= 1 && choice <= 4 {
		var n1, n2 int
		fmt.Printf("Enter two nos separated by comma:\n")
		fmt.Scanf("%d,%d", &n1, &n2)

		var result int
		switch choice {
		case 1:
			result = n1 + n2
		case 2:
			result = n1 - n2
		case 3:
			result = n1 * n2
		case 4:
			result = n1 / n2
		default:
			fmt.Println("Should not come here")
		}
		fmt.Printf("result is: %d\n", result)

	} else if choice == 5 {
		fmt.Println("Exiting")
	} else {
		fmt.Println("Invalid")
	}

}
