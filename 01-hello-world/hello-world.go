//package declaration
package main

//import other packages
import "fmt"

//package level variables

//package init function

//main function
func main() {
	/*
		var message string
		message = "Hello World!"
		fmt.Println(message)
	*/

	/*
		var message string
		var name string
		message = "Hello"
		name = "Magesh"
	*/
	/*
		var (
			message string
			name    string
		)
		message = "Hello"
		name = "Magesh"
	*/

	/*
		var (
			message, name string
		)
		message = "Hello"
		name = "Magesh"
	*/

	/*
		var message, name string
		message = "Hello"
		name = "Magesh"
	*/

	/*
		var message string = "Hello"
		var name string = "Magesh"
	*/

	/*
		var message = "Hello"
		var name = "Magesh"
	*/

	/*
		var message, name = "Hello", "Magesh"
	*/

	//the following syntax is applicable only in a function (not at the package level)
	message, name := "Hello", "Magesh"
	greet(message, name)
}

//other functions
func greet(msg string, name string) {
	fmt.Println(msg, name)
}
