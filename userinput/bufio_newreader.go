package main

import (
	"bufio"
	"fmt"
	"os"
)

func InputReader() {

	fmt.Println("Enter Your Name:")

	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Printf("Hello, %s", name)
}

/*

-> bufio.NewReader (os.Stdin) creates a new buffered reader that reads from the standard input (os.Stdin)
-> reader.ReadString('\n) reads a line from input unit it encounters a newline character ('\n'). 

----------------------------------------------------------------------------

-> A buffered reader is a type of reader that reads data from an underlying source, such as a file or standard
   input (keyboard), and stores that data in a buffer. The buffer is a temporary storage area in memory.
   Buffered readers are commonly used to improve the efficiency of input operations.
 */
