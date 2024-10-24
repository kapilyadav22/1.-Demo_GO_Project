package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func PerformPOSTRequest() {
	todo := TODO{
		UserId:    22,
		Title:     "Kapil Yadav",
		Completed: true,
	}

	// Convert the Todo struct to JSON
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshalling : ", err)
		return
	}

	// convert json data to string
	jsonString := string(jsonData)

	// convert string data to reader
	jsonReader := strings.NewReader(jsonString)

	myURL := "https://jsonplaceholder.typicode.com/todos"
	// Post takes 3 things - url, contenttype, body
	res, err := http.Post(myURL, "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error sending post request", err)
		return
	}

	defer res.Body.Close()

	// data, _ := io.ReadAll(res.Body)
	// fmt.Println("Response : ", string(data))
	fmt.Println("Response Status: ", res.Status)

}


/*
http.Post function, third argument expects an io.Reader interface for the request body.
*/