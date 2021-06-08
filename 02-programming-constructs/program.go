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
}
