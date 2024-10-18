package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//infinite loop
	counter := 0
	for {
		counter++
		fmt.Println("The value of counter is:", counter)
		if counter == 10 {
			break
		}

	}

	//range
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Index of Number is %d, and value is %d\n", index, value)
	}

	//string
	data := "Hey You are in Delhi"
	for index, char := range data {
		fmt.Printf("Index of Number is %d, and value is %c\n", index, char)
	}
}

/*

range : simplifies looping over elements of arrays, slices, maps and strings.


*/
