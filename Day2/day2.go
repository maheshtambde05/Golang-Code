package main

import "fmt"

func main() {

	fmt.Println("hello")
	// Arithmetic Operators

	a := 15
	b := 10

	fmt.Println("a+b= ", a+b)

	fmt.Println("a-b= ", a-b)

	fmt.Println("a*b= ", a*b)
	fmt.Println("a/b = ", a/b)

	//Go Decision Making (if, if-else, Nested-if, if-else-if)

	f := 700

	if f < 1000 {
		fmt.Println("f is less than 1000")
	} else if f == 1000 {
		fmt.Println("f is equal 1000")
	} else {
		fmt.Println("f is less than 1000")
	}

	// for loop

	for i := 0; i < 4; i++ {
		fmt.Println("hello")
	}
	// for var d int = 0;d<4;d++{
	// 	fmt.Println("DDDD")
	// }

	rvariable := []string{"abc", "bcd", "cde"}

	for i, j := range rvariable {
		fmt.Println(i, j)
	}

	day := 40

	switch day {

	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("invalid Day")
	}

	// Input In go

	var name string
	var age int
	fmt.Println("enter Youar Nmae and age ")
	//fmt.Scan(&name, &age)
	fmt.Scanln(&name, &age)
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
}
