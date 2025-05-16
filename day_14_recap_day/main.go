package main

import (
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	fmt.Println("Recap of Days 1–13")

	// Variables
	name := "john"
	age := 30
	fmt.Printf("Name: %s, Age: %d\n", name, age)

	// Conditionals
	if age >= 18 {
		fmt.Println("You're an adult.")
	} else {
		fmt.Println("You're a minor.")
	}

	// Loops
	fmt.Println("Counting from 1 to 5:")
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Functions
	message := greet(name)
	fmt.Println(message)

	// Slices
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice:", numbers)

	// Maps
	person := map[string]string{"name": "Alice", "city": "Paris"}
	fmt.Println("Map:", person)
}

func greet(name string) string {
	c := cases.Title(language.English)
	return "Hello, " + c.String(name) + "!"
}
