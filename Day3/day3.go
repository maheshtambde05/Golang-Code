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

}
