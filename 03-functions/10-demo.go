//deferred functions
// note - there can be multiple deferred functions

package main

import "fmt"

func main() {
	f1()
}
func f1() {
	// open processFile
	fmt.Println("f1 Begin")
	defer func() {
		fmt.Println("f1 deferred - 1")
		// close file
	}()
	defer func() {
		fmt.Println("f1 deferred - 2")
		// close file
	}()
	defer func() {
		fmt.Println("f1 deferred - 3")
		// close file
	}()
	f2()
	fmt.Println("f1 End")
}

func f2() {
	// open processFile
	fmt.Println("f2 Begin")
	defer func() {
		fmt.Println("f2 deferred")
		// close file
	}()
	f3()
	fmt.Println("f2 End")
}
func f3() {
	// open processFile
	fmt.Println("f3 Begin")
	defer func() {
		fmt.Println("f3 deferred")
		// close file
	}()
	fmt.Println("f3 End")
}
