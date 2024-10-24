package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func PerformUPDATERequest() {
	todo := TODO{
		UserId:    22,
		Title:     "Hey I am Kapil Yadav",
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

	myURL := "https://jsonplaceholder.typicode.com/todos/1"

	//Create PUT Request
	// http.NewRequest("PUT",)
	req, err := http.NewRequest(http.MethodPut, myURL, jsonReader)
	if err != nil {
		fmt.Println("Error creating  put request", err)
		return
	}
	req.Header.Set("Content-type", "application/json")

	//Send the request
	client := http.Client{}
	res, err1 := client.Do(req)
	if err1 != nil {
		fmt.Println("Error sending put request", err1)
		return
	}
	defer res.Body.Close()
	data, _ := io.ReadAll(res.Body)
	fmt.Println("Response : ", string(data))
	fmt.Println("Response Status: ", res.Status)

}
