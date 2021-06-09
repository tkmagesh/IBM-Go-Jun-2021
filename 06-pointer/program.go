package main

import "fmt"

func main() {
	var no int = 100
	var noPtr *int = &no
	fmt.Println("Address of no => ", noPtr)
	//deferencing
	fmt.Printf("Value at %v is %d\n", noPtr, *noPtr)
	//no == *(&no)
	increment(&no)
	fmt.Println(no)
	x, y := 10, 20
	swap(&x, &y)
	fmt.Println(x, y) //= > should print 20, 10

	nos := []int{10, 20, 30}
	addValue(&nos, 40)
	fmt.Println(nos) //=> [10,20,30,40]
}

func increment(x *int) {
	*x += 1
}

func swap(xPtr, yPtr *int) {
	*xPtr, *yPtr = *yPtr, *xPtr
}

func addValue(nos *[]int, no int) {
	*nos = append(*nos, no)
}
