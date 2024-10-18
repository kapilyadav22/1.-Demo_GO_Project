package main

import "fmt"

func DemoFunction() {
	fmt.Println("We are inside demo function")
}

// func Add(a, b int) int {
// 	return a + b
// }

// func Add(a int, b int) int {
// 	return a + b
// }

func Add(a, b int) (result int) {
	result = a + b
	return
}

// func

func main() {
	fmt.Println("Hey!!! Let's Learn Functions in GO")
	DemoFunction()
	fmt.Println("Addition of Two numbers are ", Add(3, 4))
}

/*

function open brackets should start from same line
	func add()
	{
	}


	func Add(a int, b ) int {
		return a + b
	}

-> will give error


*/
