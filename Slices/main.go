package main

import "fmt"

func main() {
	fmt.Println("Slices  in GO")

	numbers := []int{1, 2, 3, 4, 5}
	numbers = append(numbers, 6, 7, 8, 9, 10)
	fmt.Println("Numbers are", numbers)
	fmt.Printf("Type of Numbers are : %T\n", numbers)
	fmt.Println("Length of Numbers is : ", len(numbers))

	fmt.Println("Slice is", numbers)
	fmt.Println("Length is", len(numbers))
	fmt.Println("Capacity is", cap(numbers))

	//length = 3, capacity = 4
	values := make([]int, 3, 4)
	values = append(values, 1) //values = [0 0 0 1 ], l = 4, c=4
	values = append(values, 5) //values = [0 0 0 1 5], l=5, c= 8, capacity gets doubled if length exceeds capacity

	fmt.Println("Slice is", values)
	fmt.Println("Length is", len(values))
	fmt.Println("Capacity is", cap(values))

	var slice = make([]string, 0)
	fmt.Println("Slice is", slice)
	fmt.Println("Length is", len(slice))
	fmt.Println("Capacity is", cap(slice))

}

/*

Key Characteristics of Slice:
------------------------------------

-> Dynamic Size : Slices can grow or shrink dynamically runtime.
-> Reference to Underlying Array: Slices are referenced to underlying array, and modifying elements in the slice
will affect the original array
-> Syntax : numbers : = [] int{1,2,3,4,5}
			name :=[]string{}
-> Use of make Function : we can use make function to create a slice with a specific length and capacity.
-> Appending Elements : append function is used to add elements to a slice.


var numbers = []string , we cant initialize slice like this
-> when we declare a slice without providing any initial values, you need to use the make function to create
the slice with specified length and capacity.

-> var numbers  = make([] string, 0) //size 0 slice

*/
