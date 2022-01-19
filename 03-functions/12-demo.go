// panic and recovery
package main

import "fmt"

func main() {

	result, err := divideClient(10, 0)
	if err != nil {
		fmt.Println(result)
	} else {
		fmt.Println("Program Exiting")
	}

}

func divideClient(x, y int) (result int, err error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			err = fmt.Errorf("Unable to divide %d by %d", x, y)
			fmt.Println(err)
		} else {
			fmt.Println("program exiting successfully")
		}
	}()
	result, _ = divide(x, y)
	return
}

// recovering from panic done best in the function encountering a panic
func divide(x, y int) (q, r int) {
	if y == 0 {
		panic("Divisor cannot be zero")
	}
	return x / y, x % y
}
