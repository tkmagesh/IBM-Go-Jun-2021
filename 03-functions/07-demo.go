package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := divide(100, 0)
	if err != nil {
		fmt.Println("Something went wrong!")
		fmt.Println(err)
		return
	}
	fmt.Println("Result = ", result)
}

func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("Invalid argument(s)")
	}
	return x / y, nil
}
