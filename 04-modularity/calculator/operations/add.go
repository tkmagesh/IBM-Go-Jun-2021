package operations

import "modularity-demo/calculator"

func Add(x, y int) int {
	calculator.IncrementOpCount()
	return x + y
}
