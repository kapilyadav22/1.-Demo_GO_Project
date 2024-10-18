//directories and packages in go are tightly coupled
//each file should have package
//package of all files inside directory should be same

package myutil

import "fmt"

func PrintUtil() {
	fmt.Println("Inside util package and utils file")
}
