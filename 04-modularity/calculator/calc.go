package calculator

import "fmt"

var operationCount int

func PrintOpCount() {
	fmt.Println("Operation Count = ", operationCount)
}

func IncrementOpCount() {
	operationCount += 1
}
