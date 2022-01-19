package main

import "fmt"

var operations = map[int]func(int, int) int{
	1: add,
	2: subtract,
	3: multiply,
	4: divide,
}

func main() {

	choice := getInput()

	if operation, exists := operations[choice]; exists {
		n1, n2 := getTwoInputs()
		fmt.Printf("result is: %d\n", operation(n1, n2))
		return
	}
	if choice == 5 {
		fmt.Println("Exiting")
	} else {
		fmt.Println("Invalid")
	}
	return

}

func add(x, y int) int {
	return x + y
}
func subtract(x, y int) int {
	return x - y
}
func multiply(x, y int) int {
	return x * y
}
func divide(x, y int) int {
	return x / y
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

func getTwoInputs() (int, int) {
	var n1, n2 int
	fmt.Printf("Enter two nos separated by comma:\n")
	fmt.Scanf("%d,%d", &n1, &n2)
	return n1, n2
}
