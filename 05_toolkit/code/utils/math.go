package utils

import "fmt"

func printNum(num int) {
	fmt.Println("Current Number", num)
}

// Adds together multiple numbers
func Add(nums ...int) int {
	total := 0
	for _, num := range nums {
		printNum(num)
		total += num
	}
	return total
}