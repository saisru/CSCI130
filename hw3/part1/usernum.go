package main

import "fmt"

func main() {
	var numOne int
	var numTwo int
	fmt.Println("finding multiplication")
	fmt.Print("Please enter first number: ")
	fmt.Scan(&numOne)
	fmt.Print("Please enter second number: ")
	fmt.Scan(&numTwo)
	fmt.Println(numOne, "*", numTwo, " = ", numOne*numTwo)
}
