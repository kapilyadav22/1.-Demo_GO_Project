package main

import "fmt"

func main() {
	name := "Kapil Kumar Yadav"
	age := 27
	height := 5.11

	//println add space between parameters
	//println prints in new line
	fmt.Println("Name:", name, "Age:", age, "Height:", height)

	fmt.Printf("Name of the person is %s\n", name)
	fmt.Printf("Age of the person is %d\n", age)
	fmt.Printf("Height of the person is %f\n", height)

	//Type
	fmt.Printf("Type of the Name is %T\n", name)
	fmt.Printf("Type of the Age is %T\n", age)
	fmt.Printf("Type of the Height is %T\n", height)

	fmt.Printf("Name: %s, Age: %d, Height: %.2f \n", name, age, height)
}

/*
   	printf function (print formatted) is used for formatted printing.
   	It allows you to control the output format by using format Specifiers (similar to C).

   -> There are multiple format specifiers

   	%d : Integer
   	%s : String
   	%T : Type of the value
   	%.3f : Float Values ( where 3 specifies the precision for floating numbers)
*/
