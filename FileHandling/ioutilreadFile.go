package main

import (
	"fmt"
	"os"
)

func ReadFromFile() {
	//Read the entire file into byte slice
	content, err := os.ReadFile("hello.txt")
	if err != nil {
		fmt.Println("Error in Reading File")
		return
	}
	fmt.Println(string(content))
}

/*
os.Readfile is convenient when we want to read the entire content of the file into memory. 
-> It may not be suitable for large files
*/
