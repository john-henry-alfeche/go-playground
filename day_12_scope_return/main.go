package main

import "fmt"

// Global variable
var globalMessage = "I'm a global variable"

func main() {
	fmt.Println("Main function start")

	// Call function with no return value
	printLocalMessage()

	// Call function with return value
	sum := add(5, 3)
	fmt.Println("Sum of 5 and 3 is:", sum)

	// Variable shadowing
	shadowExample()

	fmt.Println("Accessing global variable in main:", globalMessage)
	fmt.Println("Main function end")
}

// Function demonstrating local scope
func printLocalMessage() {
	localMessage := "I'm a local variable"
	fmt.Println(localMessage)
}

// Function with parameters and return value
func add(a int, b int) int {
	return a + b
}

// Demonstrate variable shadowing and function scope
func shadowExample() {
	globalMessage := "I'm a shadowed variable inside shadowExample"
	fmt.Println("Inside shadowExample:", globalMessage)
}
