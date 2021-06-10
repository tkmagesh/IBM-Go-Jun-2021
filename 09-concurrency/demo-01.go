package main

import (
	"fmt"
	"time"
)

func print(msg string) {
	fmt.Println(msg)
}

func main() {
	go print("Hello")
	print("World")
	//fmt.Println("End of main!")
	time.Sleep(2 * time.Second)
}
