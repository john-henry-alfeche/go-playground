package main

import "fmt"

// Define interface
type Speaker interface {
	Speak() string
}

// Define struct
type Employee struct {
	Role string
}

// Implement interface method
func (e Employee) Speak() string {
	return "I am an " + e.Role
}

func interfaceDemo() {
	var speaker Speaker = Employee{Role: "Engineer"}
	fmt.Println(speaker.Speak()) // Outputs: "I am an Engineer"
}
