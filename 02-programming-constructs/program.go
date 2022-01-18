package main

import "fmt"

func main() {

	// if
	if no := 21; no%2 == 0 { // these variable have scope of the if statement blocks
		fmt.Printf("%d", no)
	} else {
		fmt.Printf("else:%d", no)
	}
	fmt.Println("")

	// for - variation 1
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Printf("sum is : %v\n", sum)

	// for - 2 - use like while
	nSum := 1
	for nSum < 100 {
		nSum += nSum
	}
	fmt.Printf("nSum is : %v\n", nSum)

	//variation 3 - infinite loop
	x := 1
	for {
		if x > 100 {
			break
		}
		x += x
	}
	fmt.Printf("x is : %v\n", x)

	// switch construct
	/*
	 rank by score
	 score 0 - 3 -> "Terrible"
	 score 4 -7 -> "not bad"
	 score 8-9 ->Goo
	 10 -> excellent
	 otherwise -> invalide
	*/

	score := 7
	switch score {
	case 0:
	case 1:
	case 2:
	case 3:
		fmt.Println("Terrible")
	case 4:
	case 5:
	case 6:
	case 7:
		fmt.Println("Not Bad")
	case 8:
	case 9:
		fmt.Println("Good")
	case 10:
		fmt.Println("Excellent")
	default:
		fmt.Println("Invalid")
	}

	score = 10
	// multi-label case element
	switch score {
	case 0, 1, 2, 3:
		fmt.Println("Terrible")
	case 4, 5, 6, 7:
		fmt.Println("Not Bad")
	case 8, 9:
		fmt.Println("Good")
	case 10:
		fmt.Println("Excellent")
	default:
		fmt.Println("Invalid")
	}

	score = 11
	// Expression switches - The switch does not have the switch expression
	switch {
	case score >= 0 && score < 4:
		fmt.Println("Terrible")
	case score >= 4 && score < 8:
		fmt.Println("Not Bad")
	case score >= 8 && score < 9:
		fmt.Println("Good")
	case score == 10:
		fmt.Println("Excellent")
	default:
		fmt.Println("Invalid")
	}

}
