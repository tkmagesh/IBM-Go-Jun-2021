package operations

import "modularity-demo/calculator"

func Subtract(x, y int) int {
	calculator.IncrementOpCount()
	return x - y
}
