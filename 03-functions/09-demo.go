//errors - they are values
package main

import (
	"errors"
	"fmt"
)

func main() {
	result, error := divide(10, 0)
	if error == nil {
		fmt.Println(result)
	} else {
		fmt.Println(error)
	}

}

// errors should always be the last as per convention
func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("Divisor is Zero")
	}
	return x / y, nil
}
