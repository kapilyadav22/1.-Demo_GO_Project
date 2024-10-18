package main

import "fmt"

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("denominator must not be 0")
	}
	return a / b, nil
}

func divide2(a, b float64) (float64, string) {
	if b == 0 {
		return 0, "denominator must not be 0"
	}
	return a / b, "nil"
}

func main() {
	fmt.Println("Error Handling in GO")

	result, _ := divide(10, 0)
	fmt.Println("division", result)

	result2, _ := divide2(10, 0)
	fmt.Println("division ", result2)
	/*
		result, err := divide(10,0)
		if err != nil {
			fmt.Errorf("Error Found")
		} else {
			fmt.Println("division", result)
		}
	*/

}

/*

->In GO, underscore (_) is used as blank identifier. It serves as a write only variable that allows you to
discard values returned by a function or to ignore specific values when you are not interesting in using them

-> The use of _ is a convension in GO to indicate that the variable is intensionally being ignored and it's a
way to make it clear for both developer and compiler that you are not planning to use that variable.


-> If we dont use fmt.Errorf and simply return a string as an error, it will still work, but we will loose
some benefits comes with fmt.Errorf

fmt.Errorf("Denominator must not be 0") -> will give warning "error strings should not be capitalized"


*/
