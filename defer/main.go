package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Starting of the program")
	data := add(4, 5)
	defer fmt.Println("Data is :", data)
	//it will run at the end
	defer fmt.Println("Middle of the program")
	fmt.Println("End of the program")

}

/*
---------------------------------------------------------------
-> Defer is used to put the execution of current line at the end
-> If there are more defer statements, then it will execute in LIFO Order
-> we generally use differ function before funtion call, so that it will execute at the end
-> This can be useful for cleanup tasks or for ensuring that resources are released in the reverse order in which they are expected.

-> If we want to close the session/db connection, we can close it at the time of starting the session using defer, so that,
even if we forget to close the session at the end, it will not create any issue.
*/
