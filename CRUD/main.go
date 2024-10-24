package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type TODO struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func PerformGETRequest() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error Getting Data")
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error in getting Responses: ", res.Status)
		return
	}

	// data, err1 := io.ReadAll(res.Body)
	// if err1 != nil {
	// 	fmt.Println("Error in reading Response: ", res.Body)
	// 	return
	// }
	// fmt.Println("Data: ", string(data))

	var todo TODO
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error Decoding response")
		return
	}
	fmt.Println("Todo: ", todo)

}

func main() {
	fmt.Println("Let's Learn CRUD")
	// PerformGETRequest()
	// PerformPOSTRequest()
	// PerformUPDATERequest()
	PerformDELETERequest()
}

/*

Either we can read  data using io.ReadAll and then unmarshal it, or we can directly use json.NewDecoder() to put json data in go struct

*/
