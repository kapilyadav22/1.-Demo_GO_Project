package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// file, err := os.Create("hello.txt")
	// if err != nil {
	// 	fmt.Println("Error Creating file: ", err)
	// 	return
	// }
	// defer file.Close()

	// content := "Hey, This is Kapil"
	// bytes, errors := io.WriteString(file, content+"\n")
	// fmt.Println("Bumber of Bytes: ", bytes)
	// if errors != nil {
	// 	fmt.Println("Error writing inside file: ", err)
	// 	return
	// }
	// fmt.Println("File Created Successfully: ", err)

	//To Open a file
	file, err := os.Open("hello.txt")
	if err != nil {
		fmt.Println("Error Opening file: ", err)
		return
	}
	defer file.Close()

	//To Read the Content From File, we use buffer
	//Create a Buffer to read the file content
	buffer := make([]byte, 1024) //1024 is buffer size

	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error while reading from file ", err)
		}

		//Process the read content
		fmt.Println(string(buffer[:n]))
	}
	ReadFromFile()
}

/*

File Operations in Go involve working with OS and io/ioutil packages.

-------------------------------------------------------------------------------------------------------------------
-> we can use io package for writing a file. The io.WriteString function is a convenient way to write a string to a file.
-> file.Close() is important in many cases when we are done working with file. When we open a file using functions like os.Create,
os.Open, or others, we are acquiring system resources to interact with that file. Failing to close the file can lead to resources leak and
might cause issues like running out of file descriptors.
-> The io.WriteString function returns two values:
1.) The number of bytes written : It is an integer representing the number of bytes written to the file.
2.) An error: return if contains any error.
*/
