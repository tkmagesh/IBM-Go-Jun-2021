package main

import "fmt"

//hide the counter variable in such a way that it is accessible only to the count function
var counter int = 0

func count() int {
	counter += 1
	return counter
}

func main() {
	fmt.Println(count())
	fmt.Println(count())
	fmt.Println(count())
	fmt.Println(count())
}
