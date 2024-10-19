package main

import "fmt"

func main() {

	studentGrades := make(map[string]int)
	studentGrades["Kapil"] = 20
	studentGrades["Rahul"] = 21
	studentGrades["Ravi"] = 22
	studentGrades["Mohit"] = 23
	studentGrades["Ajay"] = 24
	studentGrades["Aditya"] = 25
	studentGrades["Vineet"] = 26

	fmt.Println("Marks of Kapil is :", studentGrades["Kapil"])
	fmt.Println(studentGrades)

	studentGrades["Vineet"] = 100
	fmt.Println("Marks of Vineet is :", studentGrades["Vineet"])

	delete(studentGrades, "Vineet")
	fmt.Println("Marks of Vineet is :", studentGrades["Vineet"])

	//check if the key exist
	/*
		grades, exists := studentGrades["David"]
		fmt.Println("Grades of David is :", grades)
		fmt.Println("David exists :", exists)

		Grades, Exists := studentGrades["Kapil"]
		fmt.Println("Grades of David is :", Grades)
		fmt.Println("David exists :", Exists)
	*/

	for index, value := range studentGrades {
		fmt.Printf("Key is %s and marks is %d\n", index, value)
	}

	//Declare and initiliaze map at same time
	person := map[string]int{
		"kapil": 20,
		"rahul": 70,
		"ajay":  90,
	}

	for index, value := range person {
		fmt.Printf(" Key is %s and marks is %d\n", index, value)
	}

}

/*

---------------------------------------------------------------------------------------
-> map is a ds that provides an unordered collection of key-value pairs, where each key must be unique
->if the key is not present, it will return default value in this case i.e. 0


----------------------------------------------------------------------------------------
-> The make function is used to create an empty map.
-> The existence of a key is checked using the second return value of a map lookup. (exists)
-> Iteration over the map is done using a for loop
-> in case of map, index is key and value is value, so we can use range
*/
