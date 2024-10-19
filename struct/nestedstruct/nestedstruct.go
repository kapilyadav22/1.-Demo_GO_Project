package nestedstruct

import "fmt"

type Person struct {
	FirsName string
	LastName string
	Age      int
}

type Contact struct {
	Email string
	Phone string
}

type Address struct {
	House   int
	Area    string
	State   string
	Zipcode int
}

type Employee struct {
	Person_Details  Person
	Contact_Details Contact
	Address_Details Address
}

func Nestedstruct() {
	var employee1 Employee
	employee1.Person_Details = Person{
		FirsName: "Amit",
		LastName: "Kumar",
		Age:      48,
	}
	employee1.Contact_Details = Contact{Email: "singhkapil347@gmail.com",
		Phone: "844757xxxx"}
	employee1.Address_Details = Address{House: 1691,
		Area:    "Sangam Vihar",
		State:   "Delhi",
		Zipcode: 110062}

	fmt.Println("Employee All Details are : ", employee1)
	fmt.Println("Employee Person Details are : ", employee1.Person_Details)
	fmt.Println("Employee First Name is : ", employee1.Person_Details.FirsName)
	fmt.Println("Employee Contact Details are : ", employee1.Contact_Details)
	fmt.Println("Employee Address Details are : ", employee1.Address_Details)
}
