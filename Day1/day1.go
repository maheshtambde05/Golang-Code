package main

import "fmt"

const PI = 3.14879685

func main() {
	// Print Hello world in console
	fmt.Println("hello world")

	//variable declaration
	var firstName = "UserFirstName"
	var middleName = "usermiddlename"
	var lastName = "user last name"
	fmt.Printf(firstName)
	fmt.Print("\n ", middleName)
	fmt.Println(lastName)

	var myvariable1 = "Golangcode"
	var myvariable2 = 50
	var myvariable3 = 30.48

	fmt.Printf("The type of myvariable1 is : %T, and the value is :%v\n", myvariable1, myvariable1)
	fmt.Printf("the type of myvariable2 is : %T,and value is : %v \n", myvariable2, myvariable2)
	fmt.Printf("the type of myvariable3 is :%T , and value is %v \n", myvariable3, myvariable3)
	fmt.Printf("the value is %d \n", myvariable2)
	fmt.Printf("My name is %s and I am %d years old. and emp id: %d and gender is : %s and weight is : %f \n", "Mahesh", 25, 135896329, "Male", 70.50)
	//Using short variable declaration:

	myvar1 := 56
	myvar2 := "sun"
	myvar3 := 85.69

	fmt.Printf("the value for v1 : %v and type is : %T \n", myvar1, myvar1)
	fmt.Printf("the value for v1 : %v and type is : %T \n", myvar2, myvar2)
	fmt.Printf("the value for v1 : %v and type is : %T \n", myvar3, myvar3)

	myvar11, myvar22, myvar33 := 800, "Geeks", 47.56

	fmt.Println(myvar11)
	fmt.Println(myvar22)
	fmt.Println(myvar33)

	//Constants

	const abc = "constantString"

	fmt.Println(abc)
	fmt.Println(PI)

	const Correct = true
	fmt.Println("Go rules?", Correct)

	// Data Types

	var x uint8 = 225

	fmt.Println(x, x-5)
	var y int16 = 32767

	fmt.Println(y+2, y-2)

}
