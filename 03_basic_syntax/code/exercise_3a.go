package main

import "fmt"

func main() {
	var sentence = "I love swb"

	for index, letter := range sentence {
		if index % 2 == 0 {
			fmt.Println(string(letter))
		}
	}
}