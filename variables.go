package main

import "fmt"

func Varnames() {

	//we have declared it as string, so it will check if the value is string or not
	var name string = "Kapil"
	fmt.Println(name)

	//it will check the datatype on compile time
	var num = 1
	fmt.Println(num)

	var money int = 10000
	fmt.Println(money)

	var dimension float64 = 100.00001
	fmt.Println(dimension)

	var isTrue bool = false
	fmt.Println(isTrue)

	var myname = "Kapil Yadav"
	fmt.Println(myname)

	const pi = 3.14
	fmt.Println(pi)

	//without declaring data type, we use it when we are getting data from a function, and we want to store it in variable
	person := 123
	fmt.Println(person)

	//public variable
	var Pub int = 22
	fmt.Println(Pub)

	//private variable
	var priv int = 22
	fmt.Println(priv)

}

/*
run every file under main package
-> go run main.go variables.go OR
-> go run .

------------------------------------------
Variable and Function Visibility:
-> In Go, the visibility of a variable or function outside its package is determined by the capitalization of its name
-> If the first letter of a name is uppercase, it's exported (public). If it's lowercase, its unexported (private)
 and only visible within the same package.

*/
