package main

import "fmt"

func main() {
	//if construct

	/*
		no := 21
		if no%2 == 0 {
			fmt.Println("It is an even number")
		} else {
			fmt.Println("It is an odd number")
		}
		fmt.Println("The given number is ", no)
	*/

	if no := 21; no%2 == 0 {
		//fmt.Println("It is an even number")
		fmt.Printf("The given value %d of type %T is an even number\n", no, no)
	} else {
		//fmt.Println("It is an odd number")
		fmt.Printf("The given value %d of type %T is an odd number\n", no, no)
	}
	//fmt.Println("The given number is ", no)

	//for construct
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Printf("Sum of the first 10 integers is %d\n", sum)

	//for used as a while construct
	numSum := 1
	for numSum < 100 {
		numSum += numSum
	}
	fmt.Printf("for construct used as a while construct, numSum = %d\n", numSum)

	//indefinite loop
	x := 1
	for {
		if x > 100 {
			break
		}
		x += x
	}
	fmt.Printf("x = %d\n", x)

	no2 := 40
	switch no2 % 2 {
	case 0:
		fmt.Printf("The given value %d of type %T is an even number\n", no2, no2)
	case 1:
		fmt.Printf("The given value %d of type %T is an even number\n", no2, no2)
	}

	/*
		Rank by score
		score 0 - 3 => "Terrible"
		score 4 - 7 => "Not Bad"
		score 8 - 9 => "Good"
		score 10 => "Excellent"
		otherwise => "Unknown score"
	*/
	score := 8
	/*
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
			fmt.Println("Unknown score")

		}
	*/

	/*
		switch score {
		case 0, 1, 2, 3:
			fmt.Println("Terrible")
		case 4, 5, 6, 7:
			fmt.Println("Not bad")
		case 8, 9:
			fmt.Println("Good")
		case 10:
			fmt.Println("Excellent")
		default:
			fmt.Println("Unknown score")
		}
	*/
	switch {
	case score >= 0 && score <= 3:
		fmt.Println("Terrible")
	case score >= 4 && score <= 7:
		fmt.Println("Not bad")
	case score >= 8 && score <= 9:
		fmt.Println("Good")
	case score == 10:
		fmt.Println("Excellent")
	default:
		fmt.Println("Unknown score")
	}

	n := 4
	switch n {
	case 0:
		fmt.Println("is zero")
	case 1:
		fmt.Println("n <= 1")
		fallthrough
	case 2:
		fmt.Println("n <= 2")
		fallthrough
	case 3:
		fmt.Println("n <= 3")
		fallthrough
	case 4:
		fmt.Println("n <= 4")
		fallthrough
	case 5:
		fmt.Println("n <= 5")
		fallthrough
	case 6:
		fmt.Println("n <= 6")
		fallthrough
	case 7:
		fmt.Println("n <= 7")
		fallthrough
	case 8:
		fmt.Println("n <= 8")
		fallthrough
	case 9:
		fmt.Println("n <= 9")
		fallthrough
	case 10:
		fmt.Println("n <= 10")
	}

}
