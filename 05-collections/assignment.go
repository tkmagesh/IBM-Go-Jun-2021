package main

import (
	"fmt"
	"math/rand"
)

func main() {
	randomNos := genRandomNos(20)
	/*
		evenNos := filterEven(randomNos)
		primeNos := filterPrime(randomNos)
	*/
	evenNos := filter(randomNos, isEven)
	primeNos := filter(randomNos, isPrime)
	fmt.Println("Initial list => ", randomNos)
	fmt.Println("Even Nos => ", evenNos)
	fmt.Println("Prime Nos => ", primeNos)
	fmt.Println("Odd Nos => ", filter(randomNos, isOdd))
}

func genRandomNos(size int) []int {
	randomNos := make([]int, size)
	for idx := 0; idx < size; idx++ {
		randomNos[idx] = rand.Intn(100)
	}
	return randomNos
}

/*
func filterEven(nos []int) []int {
	result := []int{}
	for _, no := range nos {
		if isEven(no) {
			result = append(result, no)
		}
	}
	return result
}

func filterPrime(nos []int) []int {
	result := []int{}
	for _, no := range nos {
		if isPrime(no) {
			result = append(result, no)
		}
	}
	return result
}
*/

func filter(nos []int, predicate func(int) bool) []int {
	result := []int{}
	for _, no := range nos {
		if predicate(no) {
			result = append(result, no)
		}
	}
	return result
}

func isEven(no int) bool {
	return no%2 == 0
}

func isPrime(no int) bool {
	if no <= 2 {
		return true
	}
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}

func isOdd(no int) bool {
	return no%2 == 1
}
