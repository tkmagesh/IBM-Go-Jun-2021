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
	wg.Add(2)
	go print("Hello")
	go print("World")
	//time.Sleep(2 * time.Second)
	wg.Wait() //wait for the internal counter to become 0
	fmt.Println("End of main!")
}
