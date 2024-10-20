package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "apple,orange,bananas"
	parts := strings.Split(data, ",")
	fmt.Println(parts)

	str := "one one two two three three four four four four"
	//it is case sensitive
	count1 := strings.Count(str, "four")
	count2 := strings.Count(str, "Four")
	fmt.Println(count1)
	fmt.Println(count2)

	str = "      Hello World     "
	trimmed_data := strings.TrimSpace(str)
	fmt.Println(trimmed_data)

	firstName := "Kapil"
	lastName := "Yadav"

	//second parameter in join is separator between the firstname and lastname
	name := strings.Join([]string{firstName, lastName}, " ")
	fmt.Println(name)
}

/*
---------------------------------------------------------------
Strings : The String package in Go provides a variety of functions for manipulating strings.

Splitting Strings:


*/
