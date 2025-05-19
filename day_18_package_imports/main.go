package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	// Using the math package
	number := 64
	squareRot := math.Sqrt(float64(number))
	fmt.Printf("Square root of %d is: %f \n", number, squareRot)

	// Using the strings package
	original := "Hello, Go lang"
	upper := strings.ToUpper(original)
	fmt.Println("Original: ", original)
	fmt.Println("Uppercase: ", upper)

	// Using a custom package (same file for now)
	result := Add(10, 20)
	fmt.Println("10 + 20 = ", result)
}

// Add Simulating a custom package function
func Add(a, b int) int {
	return a + b
}
