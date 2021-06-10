package main

import "fmt"

type Employee struct {
	Name string
}

func main() {
	var x interface{}
	x = 100
	x = "Magesh"
	x = true
	fmt.Println(x)

	/*
		var z int32 = 100
		y := int64(z)
	*/

	if no, ok := x.(int); ok {
		fmt.Println("Double of x is ", no*2)
	} else {
		fmt.Println("x is not numeric")
	}

	switch val := x.(type) {
	case int:
		fmt.Println("Double of x is ", val*2)
	case string:
		fmt.Println("Len of x is ", len(val))
	case Employee:
		fmt.Println("Name of the employee is ", val.Name)
	default:
		fmt.Println("Unknown type")
	}
}
