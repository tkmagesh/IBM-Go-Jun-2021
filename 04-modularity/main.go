package main

import (
	calc "modularity-demo/calculator"
	"modularity-demo/calculator/operations"

	"github.com/fatih/color"
)

func main() {
	color.Green("Adding 100 and 200 = %d\n", operations.Add(100, 200))
	color.Red("Subtracting 100 and 200 = %d\n", operations.Subtract(100, 200))
	color.Green("Adding 10 and 20 = %d\n", operations.Add(10, 20))
	color.Red("Subtracting 10 and 20 = %d\n", operations.Subtract(10, 20))
	calc.PrintOpCount()
}
