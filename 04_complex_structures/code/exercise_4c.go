package main

import "fmt"

// Part 1
func average(scores [5]float64) float64 {
	total := 0.0
	for _, num := range scores {
		total += num
	}

	return total / float64(len(scores))
}

// Part 2
var initialPets map[string]string = map[string]string{
	"fido": 	"dog", 
	"penelope": "horse", 
	"nancy": 	"cat",
}

func doesPetExists(petName string) bool {
	_, exist := initialPets[petName]
	return exist
}

// Part 3
var initialGroceries = []string{"beans", "lemons", "chicken", "fruit"}

func addToGroceryToList(newGroceries ...string) []string {
	foods := initialGroceries
	for _, groceries := range newGroceries {
		foods = append(foods, groceries)
	}
	return foods
}

func main() {
	// scores := [5]float64{2, 7, 8, 1, 9}
	// fmt.Println(average(scores))

	// pet := "fido"
	// petExists := doesPetExists(pet)
	// fmt.Println(petExists)

	groceryList := addToGroceryToList("ass", "beets")
	fmt.Println(groceryList)
}