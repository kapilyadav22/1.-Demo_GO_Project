package main

import (
	"demogolearning/struct/nestedstruct"
	"fmt"
)

type Person struct {
	FirsName string
	LastName string
	Age      int
}

func main() {

	var kapil Person
	fmt.Println("Kapil Person : ", kapil) // "" "" 0
	kapil.FirsName = "Kapil"
	kapil.LastName = "Yadav"
	kapil.Age = 27

	fmt.Println("Kapil Person : ", kapil)

	// -----------------------------------------------------------------------------------------------------------------------------

	person1 := Person{"Ajay", "Rathore", 26}
	fmt.Println("person1 Person : ", person1)

	// -----------------------------------------------------------------------------------------------------------------------------

	person2 := Person{
		FirsName: "Amit",
		LastName: "Kumar",
		Age:      48,
	}
	fmt.Println("person2 Person : ", person2)

	// -----------------------------------------------------------------------------------------------------------------------------

	var person3 = new(Person)
	person3.FirsName = "Vineet"
	person3.LastName = "Tomar"

	fmt.Println("Vineet Information : ", person3)
	fmt.Println("Vineet FirstName : ", person3.FirsName)
	fmt.Println("Vineet LastName : ", person3.LastName)
	//age will print default value of age data type, int in our case
	fmt.Println("Vineet Age : ", person3.Age)

	nestedstruct.Nestedstruct()
}

/*
----------------------------------------------------------------------
-> struct (short of structure) is a composite data type that groups together variables (fields or members) under a single name.
-> Each field in a struct can have a diffferent data type, and structs are used to create more complex ds.


*/
