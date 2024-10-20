package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Let's Learn Type Conversion")

	var data int = 5
	fmt.Printf("data is of type %T\n", data)
	var newData float32 = float32(data)
	fmt.Printf("New data is of type %T\n", newData)

	num := 23
	str := strconv.Itoa(num)
	fmt.Printf("New data is of type %T\n", str)

	number_string := "22101997"
	//Atoi returns two values, first one is value, second one is error, we want to ignore the error, so use _
	num_int, _ := strconv.Atoi(number_string)
	num_float, _ := strconv.ParseFloat(number_string, 64)
	num_bool, _ := strconv.ParseBool("True")
	num_int2, _ := strconv.ParseInt(number_string, 10, 64)  //base 10 - decimal,8 - for octal, 16 for hex
	num_uint, _ := strconv.ParseUint(number_string, 10, 64) //base 10 -decimal,8 - for octal, 16 for hex
	fmt.Printf("number_string is of Type %T\n", number_string)
	fmt.Printf("num_int is of type %T\n", num_int)
	fmt.Printf("num_float is of type %T\n", num_float)
	fmt.Printf("num_bool is of type %T\n", num_bool)
	fmt.Printf("num_int is of type %T\n", num_int2)
	fmt.Printf("num_uint is of type %T\n", num_uint)

}

/*
Data Conversion: Converts data of one type to another

Numeric type conversion :

*/
