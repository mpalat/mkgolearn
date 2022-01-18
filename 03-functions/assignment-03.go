package main

import "fmt"

func main() {

	choice := getInput()
	if shouldExit(choice) {
		return
	}

	n1, n2 := getTwoInputs()
	fmt.Printf("result is: %d\n", calculate(choice, n1, n2))

}

func getInput() int {
	fmt.Println("1.Add")
	fmt.Println("2.Subtract")
	fmt.Println("3.Multiply")
	fmt.Println("4.Divide")
	fmt.Println("5.Exit")
	fmt.Println("Enter a Choice:")

	var choice int
	fmt.Scanf("%d\n", &choice)
	return choice
}

func shouldExit(choice int) bool {
	if choice >= 1 && choice < 5 {
		return false
	}
	if choice == 5 {
		fmt.Println("Exiting")
	} else {
		fmt.Println("Invalid")
	}
	return true
}

func getTwoInputs() (int, int) {
	var n1, n2 int
	fmt.Printf("Enter two nos separated by comma:\n")
	fmt.Scanf("%d,%d", &n1, &n2)
	return n1, n2
}

func calculate(choice, n1, n2 int) int {
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
	}
	return result
}
