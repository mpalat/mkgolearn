package main

import "fmt"

type Mystr string // creating an alias for string
func main() {

	s := Mystr("This is a sample String") // converting string into MyStr
	fmt.Println(s)
	fmt.Println(s.Length())
}
func (s Mystr) Length() int {
	return len(s)
}
