package main

import (
	"fmt"
)

func main() {
	// Arrays
	numbers := [3]int{10, 20, 30}
	fmt.Println("Array:", numbers)
	fmt.Println("First element:", numbers[0])

	// Modify an element
	numbers[1] = 25
	fmt.Println("Modified Array:", numbers)

	// Slices
	scores := []int{90, 85, 88}
	fmt.Println("Slice:", scores)
	fmt.Println("Length:", len(scores))

	// Append
	scores = append(scores, 92)
	fmt.Println("After append:", scores)

	// Slicing a slice
	subset := scores[1:3]
	fmt.Println("Sliced subset:", subset)

	// Using make to create a slice
	colors := make([]string, 3)
	colors[0] = "red"
	colors[1] = "green"
	colors[2] = "blue"
	fmt.Println("Colors slice:", colors)

	// Range over slice
	for i, score := range scores {
		fmt.Printf("Index %d: %d\n", i, score)
	}
}
