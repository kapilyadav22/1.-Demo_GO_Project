package main

import "fmt"

func main() {

	day := 5

	switch day {
	case 1:
		fmt.Println("Sunday")
	case 2:
		fmt.Println("Monday")
	case 3:
		fmt.Println("Tuesday")
	case 4:
		fmt.Println("Wednesday")
	case 5:
		fmt.Println("Thursday")
	case 6:
		fmt.Println("Friday")
	case 7:
		fmt.Println("Saturday")
	default:
		fmt.Println("Not a Valid Day")
	}

	//Switch with multiple values
	month := "January"

	switch month {
	case "January", "February", "March":
		fmt.Println("Winters")
	case "April", "May", "June":
		fmt.Println("Summers")
	case "July", "August", "September":
		fmt.Println("Monsoon")
	case "October", "November", "December":
		fmt.Println("Spring")
	}

	//Switch with Expressions
	temperature := 25

	switch {
	case temperature < 0:
		fmt.Println("Cold")
	case temperature > 100:
		fmt.Println("hot")
	case temperature > 20 && temperature < 100:
		fmt.Println("Awesome")

	}

}
