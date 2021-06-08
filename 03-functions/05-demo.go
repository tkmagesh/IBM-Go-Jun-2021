package main

import "fmt"

//hide the counter variable in such a way that it is accessible only to the count function

func getCount() (func() int, func() int) { // step 1
	var counter int = 0 //step 2

	increment := func() int { //step 3
		counter += 1 //step 4
		return counter
	}

	decrement := func() int {
		counter -= 1
		return counter
	}

	return increment, decrement // step 5
}

func main() {
	increment, decrement := getCount()
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(decrement())
	fmt.Println(decrement())
	fmt.Println(decrement())
	fmt.Println(decrement())
}
