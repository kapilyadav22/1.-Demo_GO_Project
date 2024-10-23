package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	//if we dont give json keys, then by default our variable name in struct will be our key
	Name   string `json:"name"`
	Age    int    `json:age`
	Gender string `json:"gender`
}

func main() {
	person := Person{Name: "kapil", Age: 27, Gender: "Male"}
	fmt.Println("Person Details are : ", person)

	//Marshal the data
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling ", err)
		return
	}
	fmt.Println("Person Data in JSON is :", string(jsonData))

	//UnMarshal the data
	var decodedData Person
	err1 := json.Unmarshal(jsonData, &decodedData)
	if err1 != nil {
		fmt.Println("Error Unmarshalling ", err)
		return
	}
	fmt.Println("Person Data  is :", decodedData)

}

/*
-> In Go, the encoding/json package is used to encode and decode JSON (Javascript Object Notation) data. JSON is a lightweight data-interchange format
that is easy for humans to read and write and easy for machines to parse and generate.

-> Defining a Struct : Define a struct that represents the JSON data structure. Each field in the struct should have a tag specifying the JSON key associated with it.
-> Marshalling (Encoding) : use json.Marshal to convert a GO struct into a JSON encoded byte array.
-> Unmarshalling (Decoding) : use json.Unmarshal to convert a JSON-encoded byte array into a GO struct.

*/
