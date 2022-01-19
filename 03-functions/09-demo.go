//errors - they are values
package main

import (
	"errors"
	"fmt"
)

var DivideByZeroerror = errors.New("Divisor is Zero")

func main() {
	result, error := divide(10, 0)
	if error == nil {
		fmt.Println(result)
	} else if error == DivideByZeroerror {
		fmt.Println(error) // or any other string.. just an example .. instead of clubbing iwht else
	} else {
		fmt.Println(error)
	}

	fmt.Println("Named Retiurn Values")
	result, error = divideNamed(10, 0)
	if error != nil {
		fmt.Println(result)
	} else {
		fmt.Println(error)
	}

}

// errors should always be the last as per convention
func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, DivideByZeroerror
	}
	return x / y, nil
}

func divideNamed(x, y int) (result int, err error) {
	if y == 0 {
		err = DivideByZeroerror
		return
	}
	result = x / y
	return
}
