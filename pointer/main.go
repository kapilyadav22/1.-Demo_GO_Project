package main

import "fmt"

func modifyValueByReference(value *int) {
	*value = *value + 10
}

func main() {
	fmt.Println("Let's Learn Pointer")

	// var num int
	// num = 5
	// var ptr *int
	// ptr = &num

	num := 21
	ptr := &num

	fmt.Println("Pointer refers to address ", ptr)
	fmt.Println("Pointer's address is  ", &ptr)
	fmt.Println("Pointer holds the value", *ptr)

	var myptr *int

	if myptr == nil {
		fmt.Println("Pointer default value", myptr)
	}

	value := 10
	modifyValueByReference(&value)
	fmt.Println("value contains", value)
}

/*

Pointers in GO:
-> A Pointer is a variable that stores the memory address of another variable.
-> Pointers are used to indirectly refer to the value stored in a variable, rather than the value itself.
-> They Provide a way to work with  memory directly, which can be helful for various programming tasks, including efficient memory
management and sharing data between functions.
-> To declare a pointer, we use * symbol followed by the type of variable it will point to.We can then initialize the pointer using the
ampersane (&) operator.


*/
