package main

import "fmt"

func max(numbers ...int) int {
	var largest int
	for _, v := range numbers {
		if v > largest {
			largest = v
		}
	}
	return largest
}

func main() {
	greatest := max(1, 2, 2, 7, 76, 23, 45, 3, 15)
	fmt.Println(greatest)
}
