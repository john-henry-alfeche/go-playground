package main

import (
	"fmt"
)

func main() {
	// Basic if statement
	age := 20
	if age >= 18 {
		fmt.Println("You are eligible to vote.")
	}

	// if-else statement
	score := 65
	if score >= 75 {
		fmt.Println("You passed with distinction!")
	} else {
		fmt.Println("You passed, but no distinction.")
	}

	// if-else if-else chain
	temp := 30
	if temp < 10 {
		fmt.Println("It's very cold.")
	} else if temp < 25 {
		fmt.Println("It's a nice day.")
	} else {
		fmt.Println("It's hot outside.")
	}

	// Using == and != operators
	name := "Go Explorer"
	if name == "Go Explorer" {
		fmt.Println("Welcome back, adventurer!")
	} else {
		fmt.Println("Hello, stranger!")
	}

	// Logical operators
	isMember := true
	spending := 120.0
	if isMember && spending > 100 {
		fmt.Println("You get a 10% discount.")
	} else {
		fmt.Println("No discount for you.")
	}

	// Nested conditionals
	hour := 14
	if hour >= 0 && hour < 24 {
		if hour < 12 {
			fmt.Println("Good morning!")
		} else if hour < 18 {
			fmt.Println("Good afternoon!")
		} else {
			fmt.Println("Good evening!")
		}
	} else {
		fmt.Println("Invalid hour.")
	}

	isRaining := true
	// Negation with !	isRaining := false
	if !isRaining {
		fmt.Println("Go for a walk!")
	} else {
		fmt.Println("Take an umbrella.")
	}
}
