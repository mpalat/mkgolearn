// panic and recovery
package main

import "fmt"

func main() {

	fmt.Println(divide(10, 0))

}

// recovering from panic done best in the function encountering a panic
func divide(x, y int) (result int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from panic")
			fmt.Println(err)
		} else {
			fmt.Println("program exiting successfully")
		}
	}()
	return x / y
}
