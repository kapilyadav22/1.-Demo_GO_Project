package main

import (
	"fmt"
)

func main() {
	fmt.Println("First GoLang Project")
	// myutil.PrintUtil()
	Varnames()
}

/*
-> In Go, for a file to be the entry point of an executable program, it must belong to the main package and
contain a main function.
-> When we use the go run command to execute a go program, the GO runtime looks for the main function  in
the main package and execute it.
-> We can give any file name other than main.go. It will check for main package and main function
 and execute it.
*/
