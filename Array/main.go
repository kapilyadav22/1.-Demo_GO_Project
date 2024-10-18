package main

import "fmt"

func main() {
	fmt.Println("Array  in GO")

	var name [5]string
	name[0] = "Kapil"
	name[1] = "Rahul"
	name[2] = "Ajay"

	fmt.Println("Name of first person is", name[0])
	fmt.Println("Name of  person is", name)

	var num = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Numbers are", num)

	fmt.Println("Length of Array is ", len(num))

	var names [5]string
	fmt.Printf("names of Array is %q\n", names)

	var values [5]int
	fmt.Println("names of Array is \n", values)
}

/*

When we declare an array or slice, the elements are initialized to their zero value. The zero value is the default
value assigned to the variables of certain type when no explicit value is provided.

for numeric type (int, float..) the zero value is 0.
for strings, it is an empty string "".
for boolean, it is false
for pointers or complex types, it is nil
*/
