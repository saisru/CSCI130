package main
//Assignment 2- Printing the type of variables
import (
	"fmt"
)

func main() {
	var num int = 10
	var str string = "Jeewaka"
	var b = true

	fmt.Printf("%T \n", num)
	fmt.Printf("%T \n", str)
	fmt.Printf("%T \n", b)

}

