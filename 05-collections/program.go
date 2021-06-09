package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//Array

	//var nos [5]int
	//var nos [5]int = [5]int{4, 6, 1, 3, 2}
	var nos [5]int = [...]int{4, 6, 1, 3, 2}
	fmt.Println(nos)
	fmt.Println("Array size = ", len(nos))

	//iterate an array
	/*
		for idx := 0; idx < len(nos); idx++ {
			fmt.Println(nos[idx])
		}
	*/
	//using the range construct
	/*
		for idx, no := range nos {
			fmt.Printf("nos[%d] = %d\n", idx, no)
		}
	*/
	for _, no := range nos {
		fmt.Println(no)
	}

	//to convert an array into a slice

	/*
		newNos := nos[:] //=> newNos is a slice
		newNos[0] = 100
		fmt.Println(newNos, nos)

		newNos = append(newNos, 1000)
		fmt.Println(newNos, nos)

		newNos[0] = 200
		fmt.Println(newNos, nos)
	*/

	//creating a copy of an array
	newNos := nos
	newNos[0] = 200
	fmt.Println(newNos, nos)

	//slice
	//var names []string = []string{"Magesh", "Ramesh"}

	/*
		var names []string
		names = append(names, "Magesh")
		names = append(names, "Ramesh")
	*/

	var names []string = make([]string, 2)
	names[0] = "Magesh"
	names[1] = "Ramesh"
	fmt.Println(names, len(names))

	//adding new elements
	/*
		names = append(names, "Ganesh")
		fmt.Println("After adding new element, names => ", names)
	*/
	additionalNames := []string{"Ganesh", "Rajesh", "Suresh"}
	//names = append(names, additionalNames[0], additionalNames[1])
	names = append(names, additionalNames...)
	fmt.Println("After adding new elements, names => ", names)

	//slicing
	/*
		[lo : hi] => from low to hi-1
		[lo : ] => all the elements from lo
		[ : hi] => all the elements from 0 to hi-1
		[lo : lo] => empty slice
		[lo : lo+1] => element at lo
		[:] => create a copy of the slice
	*/
	fmt.Println("names[1:3] => ", names[1:3])
	fmt.Println("names[3:] => ", names[3:])
	fmt.Println("names[:3] => ", names[:3])

	for idx := 0; idx < 10; idx++ {
		fmt.Println(rand.Intn(100))
	}

	//map
	/*
		var cityRanks map[string]int
		cityRanks["Udupi"] = 1
		cityRanks["Bengaluru"] = 5
	*/
	cityRanks := map[string]int{
		"Udupi":     1,
		"Bengaluru": 5,
	}
	fmt.Println(cityRanks)

	fmt.Println("Rank of Bengaluru => ", cityRanks["Bengaluru"])
	//add a new item
	cityRanks["Mysuru"] = 2
	fmt.Println(cityRanks)

	//removing an existing item
	delete(cityRanks, "Bengaluru")
	fmt.Println(cityRanks)

	//check if a key exists
	//cityRanks["Chennai"] = 3
	if rank, exists := cityRanks["Chennai"]; exists {
		fmt.Println("Rank of Chennai => ", rank)
	} else {
		fmt.Println("Chennai is not ranked yet!")
	}

}
