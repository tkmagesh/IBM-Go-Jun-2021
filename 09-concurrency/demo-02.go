package main

import (
	"fmt"
	"sync"
)

var wg *sync.WaitGroup = &sync.WaitGroup{}

func print(msg string) {
	fmt.Println(msg)
	wg.Done()
}

func main() {
	//wg.Add(2)
	wg.Add(1)
	go print("Hello")
	wg.Add(1)
	go print("World")
	wg.Wait() //wait for the internal counter to become 0
	fmt.Println("End of main!")
}
