package main

import "fmt"

type Employee struct {
	Name string
}

func (emp Employee) WhoAmI() {
	fmt.Println("I am ", emp.Name)
}

type Device struct {
}

func (dev Device) WhoAmI() {
	fmt.Println("I am a device")
}

type FullTimeEmployee struct {
	Employee
	Benefits string
	//Device
}

type FavouriteEmployee Employee

func main() {
	fte := FullTimeEmployee{Employee{"Magesh"}, "Healthcare"}
	fte.WhoAmI()

	favEmp := FavouriteEmployee{"Suresh"}
	fmt.Println(favEmp.Name)
}
