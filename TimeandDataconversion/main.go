package main

import (
	"fmt"
	"time"
)

func main() {

	currentTime := time.Now()
	fmt.Println("Current time: ", currentTime)
	fmt.Printf("type of currentTime is :  %T\n", currentTime)

	formatted := currentTime.Format("2006/01/02, Monday")
	fmt.Println("Formatted time is :  \n", formatted)

	formatted_text := currentTime.Format("2006/01/02, 15:04:05")
	fmt.Println("Formatted time is :  \n", formatted_text)

	//formatting fix for 12 hour format, we need to put it in 3:04 PM
	formatted_text2 := currentTime.Format("2006/01/02, 3:04 PM")
	fmt.Println("Formatted time is :  \n", formatted_text2)

	layout_string := "2006-01-02"
	dateString := "2024-10-22"
	formatted_time, _ := time.Parse(layout_string, dateString)
	fmt.Println("Formatted time is :  \n", formatted_time)

	//layout string should be in same format as our dateString, and the date month year should be date - 02, month - 01, year - 2006
	layout_string1 := "02/01/2006"
	dateString1 := "22/10/2024"
	formatted_time1, _ := time.Parse(layout_string1, dateString1)
	fmt.Println("Formatted time is :  \n", formatted_time1)

	//add 1 more day in currentTime
	new_date := currentTime.Add(24 * time.Hour)
	fmt.Println("New Date Time: ", new_date)
	formatted_new_date := new_date.Format("2006/01/02")
	fmt.Println(formatted_new_date)

}

/*
Time and Date Conversion:
"2006-01-02 15:04:05"
"year-dd-mm hr:minutes:seconds"  //hr in 24 hour format
-> Go Provides a powerful time package for handling time and date-related operations.
-> Go have time format like "2024-10-22 15:04:19"

*/
