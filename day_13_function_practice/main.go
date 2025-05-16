package main

import (
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Function to add two numbers
func add(a int, b int) int {
	return a + b
}

// Function to greet a user by name
func greet(name string) string {
	caser := cases.Title(language.English)
	return "Hello, " + caser.String(name) + "!"
}

// Function that returns both the sum and product of two numbers
func sumAndProduct(x, y int) (int, int) {
	sum := x + y
	product := x * y
	return sum, product
}

// Function to find the longest word in a slice
func findLongest(words []string) string {
	longest := ""
	for _, word := range words {
		if len(word) > len(longest) {
			longest = word
		}
	}
	return longest
}

func main() {
	// 1. Add two numbers
	fmt.Println("Sum of 5 and 3:", add(5, 3))

	// 2. Greet a name
	fmt.Println(greet("john"))

	// 3. Return multiple values (sum and product)
	sum, product := sumAndProduct(4, 5)
	fmt.Printf("Sum: %d, Product: %d\n", sum, product)

	// 4. Find the longest word
	words := []string{"apple", "banana", "cherry", "strawberry"}
	fmt.Println("Longest word:", findLongest(words))
}
