package main

import "fmt"

func multply(a, b int) int {
	a = a * 2
	return a * b
}

type Address struct {
	Name    string
	city    string
	pincode int
}

type Employee struct {
	firstName, lastName string
	age, salary         int
}

func main() {
	fmt.Println("hello Day 3")

	//fmt.Println(multply(5, 6))

	var a Address
	fmt.Println(a.city, a.pincode, a.Name)
	a1 := Address{"akshay", "pune", 412350}
	fmt.Println(a1)
	a2 := Address{"ram", "goa", 53888}
	fmt.Println(a2)

	a3 := Address{Name: "sham", city: "mumbai"}
	fmt.Println(a3)

	//nested struct

	type person struct {
		fName, lname string
		age          int
		Address      Address
	}

	P := person{
		fName: "ram", lname: "sham", age: 55, Address: Address{city: "pune", pincode: 856452},
	}
	fmt.Println(P)

	fmt.Println(P.Address)
	fmt.Println(P.Address.pincode)
	fmt.Println(P.lname)
}
