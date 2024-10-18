package main

import "fmt"

func main() {

	fmt.Println("Enter Your Name:")
	var name string

	//scan reads until the first whitespace character.
	fmt.Scan(&name)
	fmt.Println("Hello ", name)

	//To read whole line, use bufio package's NewScanner or ReadString functions
	InputReader()
}
